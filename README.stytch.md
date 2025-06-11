
# AI Ops Assistant

**AI Ops Assistant** is a backend platform that simulates intelligent internal productivity tools—like those used by Engineering, HR, and Operations teams—to automate workflows and improve operational efficiency. It’s designed as a hands-on demonstration of platform engineering principles: secure APIs, asynchronous processing, developer-focused architecture, and infrastructure automation.

This project mirrors the type of systems built by internal enablement teams at companies like Stytch—where reliability, security, and developer experience are top priorities.

---

## 🔧 Tech Stack

- **Language:** Go (1.21)
- **API Layer:** [graphql-go](https://github.com/graphql-go/graphql)
- **Database:** PostgreSQL via GORM
- **Authentication:** JWT (login, signup, `me` query)
- **Async Processing:** Background workers (Go routines)
- **Containerization:** Docker + Docker Compose
- **Infra-as-Code:** Terraform (planned)
- **CI/CD:** GitHub Actions (planned)
- **Frontend (planned):** React

---

## 🚀 Key Features

### 🧠 Intelligent Automation
Simulates AI-assisted workflows:
- Log summarization
- Ticket triage
- Changelog generation  
Each task is handled by background workers and exposed via GraphQL.

### 🔒 Secure Internal APIs
JWT-secured API endpoints support multi-role access (admin/user), with bcrypt-hashed passwords and protected routes—similar to how internal dashboards might be built for finance or HR systems.

### 🧰 Developer-Focused Architecture
Modular, testable Go code with clearly defined separation of concerns (`cmd/api`, `cmd/worker`, `internal/`) and a scalable, extensible GraphQL schema.

### 🛠 Designed for Scale
- Background job orchestration
- Docker-based local development
- Ready for cloud deployment with Terraform and CI/CD integration

---

## 📁 Project Structure

```
cmd/
  api/         # Main GraphQL API entrypoint
  worker/      # Entry points for async workers (summarizer, triage)

internal/
  auth/        # JWT and bcrypt logic
  db/          # DB initialization
  models/      # GORM models
  schema/      # GraphQL types & resolvers
  summarizer/  # Log summarization logic
  triage/      # Ticket triage logic
```

---

## ✅ MVP Features

1. **Log Summarization** – Accepts raw logs, summarizes via background job, stores results
2. **Ticket Triage** – Classifies tickets with labels like `bug`, `urgent`, etc.
3. **Changelog Generation** – Accepts structured entries and produces human-readable changelogs
4. **Secure Admin API** – Login/signup endpoints with JWT and protected GraphQL mutations

---

## 🗺️ Architecture Overview

![Architecture](architecture_diagram.png)

- Go-based GraphQL API handles secure CRUD operations
- Background workers simulate async task processing
- PostgreSQL stores task metadata and user info
- Docker Compose supports local development
- Terraform planned for AWS provisioning

---

## 🔍 Sample GraphQL Mutation

```graphql
mutation {
  summarizeLog(raw: "Disk failure in node-3 at 2:12am") {
    id
    summary
  }
}
```

---

## 📌 Future Plans

- [ ] GitHub Actions pipeline
- [ ] Terraform AWS deployment
- [ ] React-based internal dashboard
- [ ] Observability/metrics endpoints

---

## 🧩 Relevance

This project is intended to reflect:
- Secure, production-grade backend patterns
- Internal developer enablement tools
- Practical Go architecture with extensibility in mind
- Real-world API and infrastructure trade-offs (DX, security, and reliability)

---

## 📄 License

MIT © 2025

---

## 📚 See Also

- [System Design One-Pager](./AI_Ops_Assistant_System_Design.md)
