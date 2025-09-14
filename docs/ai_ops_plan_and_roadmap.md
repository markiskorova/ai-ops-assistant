## 🧠 AI Ops Assistant – Project Plan & Roadmap (Merged V7)

---

### 1. 📌 Purpose & Vision

AI Ops Assistant simulates a backend platform for internal operational workflows (log summarization, ticket triage, changelog generation). It demonstrates backend leadership, observability, infrastructure fluency, and secure SaaS-like patterns.

---

### 2. 🧩 Key Concepts to Integrate

1. Schema Evolution & Backward Compatibility  
2. Data Versioning & Deltas  
3. Concurrency & Invariants  
4. System Design Fundamentals  
5. Operational Safety & Deploys  
6. **Observability: Metrics, Logging, Dashboards**  
7. API Design & Evolution  
8. Security & Access Control  
9. Reliability & Failure Recovery  
10. Infrastructure & Automation  

---

### 3. 🧱 Architecture & Tech Stack

- **Backend:** Go + graphql-go, PostgreSQL, GORM  
- **AI/NLP:** OpenAI API (summarization)  
- **Auth:** JWT + bcrypt  
- **Infrastructure:** Docker Compose (local), Terraform (AWS ECS, RDS, IAM, Secrets)  
- **DevOps:** GitHub Actions CI/CD  
- **Observability:** Prometheus, Grafana, structured logs, trace IDs  
- **Frontend (Planned):** React dashboard  

---

### 4. 📆 Phase Roadmap

#### ✅ Phase 1 – Planning & Setup (Complete)
- Define MVP functionality
- Select tech stack: Go, PostgreSQL, GraphQL, JWT, Docker, Terraform
- Initialize GitHub repo and core structure: `cmd/api`, `cmd/worker`, `internal/`
- Setup Docker Compose and database scaffolding
- Scaffold Terraform for AWS deployment

**Covers:** Project bootstrapping and dev environment setup

---

#### ✅ Phase 2 – MVP Core Functionality (Complete)
- Log summarization (mocked → OpenAI integration)
- Ticket triage and classification
- Changelog generation from commit-like input
- User auth: signup/login with JWT and bcrypt
- Background worker setup and async queue
- Manual testing and `go test ./...`

**Covers:** Initial APIs, async design, security/auth foundations

---

#### 🔧 Phase 3 – Infrastructure & Observability (In Progress)
- Provision AWS: ECS Fargate, RDS, IAM, Secrets Manager
- Enable GitHub Actions CI/CD: build, test, deploy
- **Observability Enhancements:**
  - Expose Prometheus metrics from API & worker (`/metrics` endpoint)
  - Add job metrics: durations, failures, queue depth
  - Deploy exporters: Postgres exporter, node_exporter, cAdvisor
  - Add Prometheus server with alert rules (latency, error rate, backlog)
  - Deploy Grafana with provisioned dashboards (API latency/error %, worker success/failure, DB, container CPU/mem)
- Add structured logging (zap/logrus) + trace IDs
- Implement retries with backoff + jitter, simulate circuit breaker + DLQ

**Topics:** Observability, Reliability, Infra & Automation

---

#### 🔁 Phase 4 – Schema Evolution & History
- Add `log_changes` and `ticket_changes` tables for versioned deltas
- Add replay queries for time-travel/history
- Use dual-write logic for renamed/modified fields
- Add `migration_versions` table and phased schema evolution strategy

**Topics:** 1. Schema Evolution, 2. Data Versioning

---

#### 🔐 Phase 5 – Security, API Design, Concurrency
- Add role-based field masking and access control (admin/ops/HR)
- Redact PII in logs
- Add pagination, validation, and idempotency tokens to GraphQL
- Implement optimistic locking for `updateStatus` mutations

**Topics:** 3. Concurrency, 7. API Design, 8. Security

---

#### 🚦 Phase 6 – Operational Safety & Deploy Strategies
- Add feature flags system to control rollout
- Simulate blue/green deployments
- Use `migration_versions` for phased schema changes
- Document rollout + rollback strategy

**Topics:** 5. Operational Safety

---

#### 💻 Phase 7 – React Dashboard (Planned UI Layer)
- Simple UI: JWT login + GraphQL queries
- Views: changelogs, logs, tickets (with role-based filtering)
- Optional: deploy to S3/CloudFront; add demo walkthrough video

**Topics:** 4. System Design, 7. API Usage

---

#### 🧠 Phase 8 – AI/NLP Innovation Layer
- Add prompt chaining and template support
- Enable optional multi-model support

**Topics:** Applied NLP & extensibility

---

### 5. 🧪 Testing & Deployment

- Unit + integration tests
- CI/CD pipeline (GitHub Actions)
- Docker healthchecks + staging/prod environments
- Validation of Prometheus configs & dashboards in CI

---

### 6. 📘 Documentation & Showcase

- Expanded README with Observability section
- `docs/interview_topics.md` linking features to backend concepts
- System design diagrams + dashboard screenshots

---

### 7. 🔭 Future Enhancements

- Advanced Grafana dashboards (team productivity KPIs)
- Real event ingestion (SNS, S3, webhooks)
- Role-extended frontend
- NLP fine-tuning or plugin-based registry

---

**Key Integration:** Prometheus + Grafana are fully part of Phase 3, with dashboards, exporters, and alerting rules forming part of the core deliverables.

