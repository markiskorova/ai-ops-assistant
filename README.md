# AI Ops Assistant

AI Ops Assistant is a backend service that automates key DevOps tasks including:

- Log triage and summarization
- Ticket triage and enrichment
- Changelog generation from commit and deployment metadata

## Features

- REST + GraphQL APIs
- JWT-based authentication
- Background processing via worker service
- Containerized with Docker
- Infrastructure as Code with Terraform

## Tech Stack

- Python / FastAPI
- PostgreSQL
- Docker & Docker Compose
- Terraform (AWS)
- GitHub Actions (CI)

## Getting Started

```bash
# Clone the repository
git clone https://github.com/markiskorova/ai-ops-assistant.git
cd ai-ops-assistant

# Start services with Docker
docker-compose up --build
```

## Project Structure

```
ai-ops-assistant/
├── cmd/             # API and worker entrypoints
├── internal/        # Core logic (auth, db, scraper, summarizer)
├── graph/           # GraphQL schema and resolvers
├── terraform/       # Infra configuration
├── docker-compose.yml
├── .github/workflows/
└── README.md
```

## License

MIT
