# 🧠 AI Ops Assistant

AI Ops Assistant is a cloud-native backend project simulating an AI-powered operational triage and summarization platform. Designed with scalability and team productivity in mind, it showcases engineering patterns applicable to internal tooling, data processing, and privacy-conscious automation.

Built using Go, GraphQL, Docker, and Terraform with secure JWT auth, this system processes logs and tickets asynchronously, summarizes them via OpenAI, and delivers insights via a robust API—making it ideal for platforms focused on data control, governance, and observability.

---

## ⚙️ Tech Stack

- **Backend:** Go 1.23, GORM (PostgreSQL), GraphQL (graphql-go)
- **AI Integration:** OpenAI API for summarization
- **Auth:** JWT
- **Infra:** Docker, Docker Compose, Terraform (AWS)
- **DevOps:** GitHub Actions (CI/CD ready)

---

## ✅ Key Features

- 🧾 Summarizes logs using OpenAI's GPT API.
- 🏷️ Classifies tickets with pluggable business logic.
- 🌐 GraphQL API for querying logs, tickets, and changelogs.
- 🔒 JWT-secured authentication and modular user management.
- 🧵 Cleanly separated microservice-style API and worker processes.
- 🏗️ Infrastructure-as-Code via Terraform (AWS RDS).
- 🐳 Local development via Docker Compose with minimal setup.

---

## 📐 System Design

See the [System Design One-Pager](./docs/AI_Ops_Assistant_System_Design.md) for an architecture breakdown.

![Architecture Diagram](./docs/architecture_diagram.png)

---
---

## 🧪 Run Locally

```bash
docker-compose up --build
```

Then access the GraphQL endpoint at:

```
http://localhost:8080/graphql
```

---

## 🛣️ Project Roadmap

See the [MVP Roadmap](./docs/ai_ops.mvp_roadmap.md) for phased implementation and future plans.

---

## 🧑‍💻 Author Notes

This project was created to demonstrate backend leadership and infrastructure fluency aligned with real-world SaaS tooling. Its design prioritizes modularity, secure data handling, and developer productivity.

---

MIT License
