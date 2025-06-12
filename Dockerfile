
FROM golang:1.23-alpine

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build all three services
RUN go build -o /bin/api ./cmd/api/main.go
RUN go build -o /bin/summarizer ./cmd/worker/summarizer/main.go
RUN go build -o /bin/triage ./cmd/worker/triage/main.go
