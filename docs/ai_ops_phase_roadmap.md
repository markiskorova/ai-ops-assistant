## 📆 AI Ops Assistant – Phase Roadmap (V6)

This roadmap builds upon the original MVP roadmap and aligns each phase with the 10 core backend engineering topics identified for interview preparation.

---

### ✅ Phase 1 – Planning & Setup (Complete)
- Define MVP functionality
- Select tech stack: Go, PostgreSQL, GraphQL, JWT, Docker, Terraform
- Initialize GitHub repo and core structure: `cmd/api`, `cmd/worker`, `internal/`
- Setup Docker Compose and database scaffolding
- Scaffold Terraform for AWS deployment

**Covers:** Project bootstrapping and dev environment setup

---

### ✅ Phase 2 – MVP Core Functionality (Complete)
- Log summarization (mocked → OpenAI integration)
- Ticket triage and classification
- Changelog generation from commit-like input
- User auth: signup/login with JWT and bcrypt
- Background worker setup and async queue
- Manual testing and `go test ./...`

**Covers:** Initial APIs, async design, security/auth foundations

---

### 🔧 Phase 3 – Infrastructure, Observability, Reliability
- Provision AWS: RDS, ECS Fargate, Secrets Manager
- Enable GitHub Actions CI/CD: build, test, deploy
- Add Prometheus metrics and trace ID injection
- Add structured logs (logrus/zap)
- Implement retries with exponential backoff + jitter
- Simulate circuit breaker pattern and DLQ (DB fallback or failure table)

**Topics:** 6. Observability, 9. Reliability, 10. Infra & Automation

---

### 🔁 Phase 4 – Schema Evolution, Deltas, History
- Add `log_changes` and `ticket_changes` tables for versioned deltas
- Add `replayChanges(log_id)` and `replayChanges(ticket_id)` queries
- Use dual-write logic for renamed/modified fields
- Add `migration_versions` table and phased schema evolution strategy

**Topics:** 1. Schema Evolution, 2. Data Versioning

---

### 🔐 Phase 5 – Security, API Design, Concurrency
- Add role-based field masking and access control (admin/ops/HR)
- Secure sensitive fields: redact PII in logs
- Add pagination, validation, and idempotency tokens to GraphQL
- Implement optimistic locking for `updateStatus` mutations

**Topics:** 3. Concurrency, 7. API Design, 8. Security

---

### 🚦 Phase 6 – Operational Safety & Deployment Strategies
- Add feature flags system to control rollout
- Simulate blue/green deployment in local/infra setup
- Use `migration_versions` for stepwise schema changes
- Document rollout + rollback strategy

**Topics:** 5. Operational Safety

---

### 💻 Phase 7 – React Dashboard (Planned UI Layer)
- Simple UI: JWT login + GraphQL queries
- Views: changelogs, logs, tickets (with role-based filtering)
- Optional: deploy to S3/CloudFront; add demo walkthrough video

**Topics:** 4. System Design, 7. API Usage

---

### 🧠 Phase 8 – AI/NLP Innovation Layer
- Support prompt templates and chaining logic
- Enable optional model selection or OpenAI tuning hooks

**Topics:** Applied NLP & extensibility

---

### 🔭 Future Enhancements
- Observability dashboards (Grafana)
- Real event ingestion (SNS, S3, webhooks)
- Extended role-based frontend
- Optional NLP model fine-tuning registry

---

This roadmap ensures every implementation decision aligns with industry concepts and strengthens both real-world and interview-readiness.

