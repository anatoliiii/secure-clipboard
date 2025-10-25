package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/anatoliiii/secure-clipboard/internal/auth"
	"github.com/anatoliiii/secure-clipboard/internal/clipboard"
	httpserver "github.com/anatoliiii/secure-clipboard/internal/http"
)

func main() {
	var (
		listenAddr     = flag.String("listen", ":8080", "HTTP listen address")
		allowedOrigins = flag.String("cors", "", "comma separated list of allowed CORS origins")
	)
	flag.Parse()

	store := clipboard.NewInMemoryStore()

	credentials := readCredentials()
	authenticator, err := auth.NewStaticAuthenticator(credentials)
	if err != nil {
		log.Fatalf("failed to configure authenticator: %v", err)
	}

	server := httpserver.NewServer(store, authenticator, splitAndTrim(*allowedOrigins))

	srv := &http.Server{
		Addr:              *listenAddr,
		Handler:           server.Handler(),
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("starting secure clipboard server on %s", *listenAddr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func readCredentials() []string {
	value := os.Getenv("SECURE_CLIPBOARD_USERS")
	if value == "" {
		log.Println("SECURE_CLIPBOARD_USERS not set, allowing only temporary demo account demo=demo")
		return []string{"demo=demo=Demo User"}
	}
	return splitAndTrim(value)
}

func splitAndTrim(value string) []string {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
