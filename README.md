# AI Ops Assistant

AI Ops Assistant is a backend system designed to simulate intelligent operations (ITOps) workflows like log summarization, ticket triage, and changelog generation. It provides a GraphQL API for internal dashboards and automation tools, and supports background processing via worker services.

---

## 🧰 Tech Stack

- **Language:** Go (1.21)
- **API Layer:** [graphql-go](https://github.com/graphql-go/graphql)
- **Database:** PostgreSQL (via GORM)
- **Auth:** JWT (login, signup, me)
- **Async Processing:** Background worker services (Go)
- **Containerization:** Docker + Docker Compose
- **Infrastructure:** Terraform (planned)
- **Frontend (planned):** React

---

## ✅ MVP Overview

The MVP includes four core features:

1. **Log Summarization**
   - Accept raw log data
   - Generate a summary (mocked NLP)
   - Store and retrieve via GraphQL
   - Async summarizer worker included

2. **Ticket Triage**
   - Accept and classify tickets
   - Store in DB
   - Async triage worker included
   - Filter by status

3. **Changelog Generation**
   - Accept commit-like entries
   - Generate structured changelogs
   - Store grouped output as JSON
   - Query by ID or list

4. **Secure Admin API**
   - JWT login and signup
   - Passwords hashed with bcrypt
   - Protected queries and mutations
   - `me` query returns user info

---

## 📦 Project Structure

```
cmd/
  api/         # GraphQL server
  worker/
    summarizer/
    triage/

internal/
  auth/        # JWT helpers
  db/          # DB connection + init
  models/      # GORM models
  schema/      # GraphQL types & resolvers
  summarizer/  # Summarization logic
  triage/      # Ticket classification logic

.env
Dockerfile
docker-compose.yml
```

---

## 🚀 Current Progress

### ✅ Completed
- Full GraphQL API (modular schema)
- Log summarization and ticket triage (API + background workers)
- Changelog generation logic + query support
- JWT login/signup + bcrypt + `me` query
- All MVP features complete

### 🛠 Planned Next
- CI/CD via GitHub Actions
- Terraform-based AWS deployment
- React dashboard
- Metrics/observability layer

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

---

## 📄 License

MIT (c) 2025