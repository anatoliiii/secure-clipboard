package http

import (
	"encoding/json"
	stdhttp "net/http"
	"strings"

	"github.com/anatoliiii/secure-clipboard/internal/auth"
	"github.com/anatoliiii/secure-clipboard/internal/clipboard"
)

// Server exposes a HTTP API for interacting with the shared clipboard.
type Server struct {
	store          clipboard.Store
	authenticator  auth.Authenticator
	allowedOrigins []string
}

// NewServer configures a new HTTP server instance.
func NewServer(store clipboard.Store, authenticator auth.Authenticator, allowedOrigins []string) *Server {
	return &Server{store: store, authenticator: authenticator, allowedOrigins: allowedOrigins}
}

// Handler returns the HTTP handler with all routes registered.
func (s *Server) Handler() stdhttp.Handler {
	mux := stdhttp.NewServeMux()

	mux.HandleFunc("/healthz", s.handleHealthz)
	mux.HandleFunc("/api/v1/clipboard/text", s.withAuth(s.handleTextClipboard))

	return s.withCORS(mux)
}

type textClipboardResponse struct {
	Owner     string `json:"owner"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}

type textClipboardRequest struct {
	Content string `json:"content"`
}

func (s *Server) handleHealthz(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	w.WriteHeader(stdhttp.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func (s *Server) handleTextClipboard(w stdhttp.ResponseWriter, r *stdhttp.Request, user auth.User) {
	switch r.Method {
	case stdhttp.MethodGet:
		s.getTextClipboard(w, r)
	case stdhttp.MethodPost:
		s.postTextClipboard(w, r, user)
	default:
		stdhttp.Error(w, "method not allowed", stdhttp.StatusMethodNotAllowed)
	}
}

func (s *Server) getTextClipboard(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	scope := scopeFromRequest(r)

	entry, err := s.store.Get(r.Context(), scope)
	if err != nil {
		if err == clipboard.ErrEmpty {
			stdhttp.Error(w, "clipboard is empty", stdhttp.StatusNotFound)
			return
		}
		stdhttp.Error(w, "failed to read clipboard", stdhttp.StatusInternalServerError)
		return
	}
	respondJSON(w, stdhttp.StatusOK, textClipboardResponse{
		Owner:     entry.Owner,
		Content:   entry.Content,
		UpdatedAt: entry.UpdatedAt.Format(stdhttp.TimeFormat),
	})
}

func (s *Server) postTextClipboard(w stdhttp.ResponseWriter, r *stdhttp.Request, user auth.User) {
	var payload textClipboardRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		stdhttp.Error(w, "invalid JSON payload", stdhttp.StatusBadRequest)
		return
	}
	if strings.TrimSpace(payload.Content) == "" {
		stdhttp.Error(w, "content must not be empty", stdhttp.StatusBadRequest)
		return
	}

	scope := scopeFromRequest(r)

	entry, err := s.store.Set(r.Context(), scope, user.Username, payload.Content)
	if err != nil {
		stdhttp.Error(w, "unable to update clipboard", stdhttp.StatusInternalServerError)
		return
	}

	respondJSON(w, stdhttp.StatusCreated, textClipboardResponse{
		Owner:     entry.Owner,
		Content:   entry.Content,
		UpdatedAt: entry.UpdatedAt.Format(stdhttp.TimeFormat),
	})
}

func (s *Server) withAuth(next func(stdhttp.ResponseWriter, *stdhttp.Request, auth.User)) stdhttp.HandlerFunc {
	return func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic realm=\"clipboard\"")
			stdhttp.Error(w, "authentication required", stdhttp.StatusUnauthorized)
			return
		}

		user, err := s.authenticator.Authenticate(r.Context(), username, password)
		if err != nil {
			stdhttp.Error(w, "authentication failed", stdhttp.StatusUnauthorized)
			return
		}

		next(w, r, user)
	}
}

func (s *Server) withCORS(next stdhttp.Handler) stdhttp.Handler {
	return stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && s.isAllowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}
		if r.Method == stdhttp.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type,Scope")
			w.WriteHeader(stdhttp.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) isAllowedOrigin(origin string) bool {
	if len(s.allowedOrigins) == 0 {
		return false
	}
	for _, allowed := range s.allowedOrigins {
		if allowed == "*" || allowed == origin {
			return true
		}
	}
	return false
}

func respondJSON(w stdhttp.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func scopeFromRequest(r *stdhttp.Request) string {
	scope := strings.TrimSpace(r.URL.Query().Get("scope"))
	if scope == "" {
		scope = strings.TrimSpace(r.Header.Get("Scope"))
	}
	if scope == "" {
		return clipboard.DefaultScope
	}
	return scope
}
