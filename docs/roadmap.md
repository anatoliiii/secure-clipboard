# Secure Clipboard Roadmap

## Product Goals
- Provide a reliable clipboard synchronization service for securely sharing snippets across devices.
- Maintain strong security guarantees while ensuring low-latency access to clipboard data.
- Offer deployment flexibility for self-hosted and managed environments.

## Key User Scenarios
1. **Personal multi-device workflow:** A single user keeps clipboard contents synchronized across desktop and laptop devices.
2. **Team collaboration:** Team members share temporary clipboard entries to speed up pair-programming or support tasks.
3. **Secure handoff:** Sensitive snippets are shared with automatic expiry and audit visibility for compliance-sensitive teams.

## Planned Stages
- **MVP – Shared Clipboard:** Implement a single shared clipboard with basic CRUD operations and minimal configuration.
- **Stage 2 – Local/Cloud Modes:** Introduce selectable operating modes for self-hosted local deployments and managed cloud instances.
- **Stage 3 – Clients & History:** Deliver native clients (desktop, CLI) with clipboard history management and search.

## Technical Decisions
- **API Requirements:**
  - RESTful endpoints for clipboard read/write/delete, authentication token management, and health checks.
  - Support for streaming updates via Server-Sent Events or WebSockets in later stages.
- **Internal Module Changes:**
  - `internal/clipboard`: extend interface to support multi-tenant contexts, history retention policies, and pluggable storage drivers.
  - `internal/http`: add authentication middleware, API versioning, and upgrade paths for streaming protocols.
- **Storage Options:**
  - In-memory store for development and MVP deployments.
  - Redis-backed implementation for scalable shared state.
  - SQLite persistence for lightweight, file-based deployments with history support.

## Readiness Criteria
- Clear API specifications documented and covered by integration tests.
- Storage drivers provide consistent behavior across supported modes with validation of failover scenarios.
- Authentication and authorization flows are defined with threat modeling and logging requirements captured.
- Client applications can successfully read, write, and fetch history entries against staging environments.

## Open Questions
- Which authentication providers (OAuth, API keys, SSO) are required for initial release?
- How will data encryption and secret management be handled for both local and cloud modes?
- What is the roadmap for mobile (iOS/Android) clients and push synchronization?

For additional context, see the project overview in the main README.
