# AI Ops Assistant

AI Ops Assistant is a backend system designed to simulate intelligent operations (ITOps) workflows like log summarization, ticket triage, and changelog generation. It provides a GraphQL API for internal dashboards and automation tools, and supports background processing via worker services.

---

## 🧰 Tech Stack

- **Language:** Go (1.21)
- **API Layer:** [graphql-go](https://github.com/graphql-go/graphql)
- **Database:** PostgreSQL (via GORM)
- **Auth:** JWT (planned)
- **Async Processing:** Background worker service (Go)
- **Containerization:** Docker + Docker Compose
- **Infrastructure:** Terraform (planned)
- **Frontend (planned):** React

---

## ✅ MVP Overview

The MVP includes four core features:

1. **Log Summarization**
   - Accept raw log data
   - Generate a summary (currently mocked)
   - Store and retrieve via GraphQL

2. **Ticket Triage**
   - Accept tickets with status and message
   - Classify/store in DB
   - Triage logic (worker-based) is planned

3. **Changelog Generation**
   - Accept commit-like entries
   - Generate release changelogs
   - Placeholder for future implementation

4. **Secure Admin API**
   - JWT-protected access (planned)
   - User login and role-based access control

---

## 📦 Project Structure

```
cmd/
  api/       # GraphQL server
  worker/    # Background summarization worker

internal/
  db/        # DB connection + init
  models/    # GORM models
  schema/    # GraphQL types & resolvers
  summarizer/# Summarization logic (mocked)

.env         # Environment variables
Dockerfile   # Shared Docker build
docker-compose.yml
```

---

## 🚀 Current Progress

### ✅ Completed
- GraphQL API with graphql-go
- LogEntry model + `summarizeLog` mutation
- `logEntry(id)` and `logEntries(limit)` queries
- Dockerized PostgreSQL
- Background summarization worker (Docker)
- Modular schema structure

### 🔜 In Progress
- Ticket triage query + async worker
- Changelog schema and resolver

### 🛠 Planned
- JWT middleware + secure user auth
- CI/CD via GitHub Actions
- Dashboard with React

---

## 🧪 Example Query

```graphql
mutation {
  summarizeLog(raw: "Server crashed at 2am with out-of-memory error.") {
    id
    summary
  }
}
```

```graphql
query {
  logEntries(limit: 5) {
    id
    summary
    created_at
  }
}
```

---

## 📄 License

MIT (c) 2025