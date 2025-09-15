## 🧠 AI Ops Assistant – Project Plan & Roadmap (Merged V8)

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
- **Observability:** Prometheus, Grafana, Alertmanager, structured logs, trace IDs  
- **Frontend (Planned):** React dashboard  

---

### 4. 📆 Phase Roadmap

#### ✅ Phase 1 – Planning & Setup (Complete)
- Defined MVP functionality  
- Selected tech stack: Go, PostgreSQL, GraphQL, JWT, Docker, Terraform  
- Initialized GitHub repo and core structure: `cmd/api`, `cmd/worker`, `internal/`  
- Setup Docker Compose and database scaffolding  
- Scaffolded Terraform for AWS deployment  

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

#### ✅ Phase 3 – Infrastructure & Observability (Complete)
- Provisioned AWS scaffolding: ECS Fargate, RDS, IAM, Secrets Manager (Terraform planned)  
- Enabled GitHub Actions CI/CD: build, test, deploy (pipeline ready)  
- **Observability Enhancements (done):**
  - Exposed Prometheus metrics from API & workers (`/metrics` endpoints)  
  - Added job metrics: durations, failures, queue depth  
  - Integrated Prometheus server with scrape configs  
  - Deployed Grafana with **provisioned dashboards** (API throughput, latency P95/P99, error %, worker throughput/failures, queue depth)  
  - Fixed provisioning via Docker Compose (volume mounts + datasource UID)  
  - Validated dashboards with Postman Runner load tests  
- **Next increments:** Add exporters (Postgres, node_exporter, cAdvisor), alert rules (latency, error rate, backlog), structured logging, retries with backoff/jitter, DLQ simulation  

**Topics:** Observability, Reliability, Infra & Automation  

---

#### 🔁 Phase 4 – Schema Evolution & History
- Add `log_changes` and `ticket_changes` tables for versioned deltas  
- Add replay queries for time-travel/history  
- Use dual-write logic for renamed/modified fields  
- Add `migration_versions` table and phased schema evolution strategy  

**Topics:** Schema Evolution, Data Versioning  

---

#### 🔐 Phase 5 – Security, API Design, Concurrency
- Add role-based field masking and access control (admin/ops/HR)  
- Redact PII in logs  
- Add pagination, validation, and idempotency tokens to GraphQL  
- Implement optimistic locking for `updateStatus` mutations  

**Topics:** Concurrency, API Design, Security  

---

#### 🚦 Phase 6 – Operational Safety & Deploy Strategies
- Add feature flags system to control rollout  
- Simulate blue/green deployments  
- Use `migration_versions` for phased schema changes  
- Document rollout + rollback strategy  

**Topics:** Operational Safety  

---

#### 💻 Phase 7 – React Dashboard (Planned UI Layer)
- Simple UI: JWT login + GraphQL queries  
- Views: changelogs, logs, tickets (with role-based filtering)  
- Optional: deploy to S3/CloudFront; add demo walkthrough video  

**Topics:** System Design, API Usage  

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
- Updated System Design One-Pager with Prometheus & Grafana  
- `docs/interview_topics.md` linking features to backend concepts  
- System design diagrams + dashboard screenshots  

---

### 7. 🔭 Future Enhancements

- Advanced Grafana dashboards (team productivity KPIs)  
- Real event ingestion (SNS, S3, webhooks)  
- Role-extended frontend  
- NLP fine-tuning or plugin-based registry  

---

**Key Integration:** Prometheus + Grafana are now fully integrated in Phase 3, with working dashboards, metrics, and infrastructure setup. Exporters + alerting rules are the next incremental observability steps.  
