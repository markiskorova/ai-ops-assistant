# AI Ops Assistant – MVP Roadmap & Status (Updated June 2025)

## 📌 MVP Overview
The MVP focuses on 4 core functionalities:
1. **Log Summarization** – Generate concise summaries of raw logs.
2. **Ticket Triage** – Classify/prioritize tickets based on severity/category.
3. **Changelog Generation** – Generate structured release changelogs.
4. **Secure Admin API** – Provide JWT-authenticated access to all data.

---

## ✅ Phase 1 – Planning & Setup (Complete)
- [x] Define MVP functionality
- [x] Select tech stack (Go, GraphQL-go, PostgreSQL, JWT, Docker, Terraform, React)
- [x] Create project structure (API, worker, internal packages)
- [x] Setup GitHub repository
- [x] Initialize Docker Compose with PostgreSQL
- [x] Setup graphql-go with modular schema layout
- [x] Initialize Terraform configuration

---

## 🚀 Phase 2 – MVP Backend Implementation

### ✅ Step 5: Log Summarization (Complete)
- [x] Create `LogEntry` PostgreSQL model
- [x] Add `summarizeLog` mutation to GraphQL schema
- [x] Store logs and mock summaries in DB
- [x] Add `logEntry(id)` query to GraphQL
- [x] Add `logEntries(limit)` query
- [x] Add `created_at` timestamp to model
- [x] Create standalone summarization worker
- [x] Run worker in Docker using shared .env
- [x] Confirm background summarization loop works

### 🟡 Step 6: Ticket Triage (Next)
- [x] Add `TriageTicket` mutation in graphql-go
- [x] Connect mutation to PostgreSQL via GORM
- [ ] Add `ticket` query to retrieve by ID
- [ ] Implement async worker-based triage
- [ ] Add ticket listing or filtered query

### 🟡 Step 7: Changelog Generation
- [ ] Add `GenerateChangelog` mutation to GraphQL schema
- [ ] Define input format and structure (e.g., commit messages or structured entries)
- [ ] Process commit-like messages into structured changelogs
- [ ] Add database model for changelogs
- [ ] Add query for retrieving changelogs

### 🟡 Step 8: Secure Admin API & Auth
- [ ] Add JWT middleware for GraphQL endpoint
- [ ] Create user table and login mutation
- [ ] Restrict mutations/queries to authorized users

---

## 📦 Planned for Future
- CI/CD pipeline via GitHub Actions
- Real log/ticket ingestion (S3, webhook, SNS)
- Extended frontend dashboard with metrics/visualizations
- Add observability (tracing/logging)
- Optional: model fine-tuning or plugin NLP enhancements