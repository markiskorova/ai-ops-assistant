# Dockerfile
FROM golang:1.22-alpine

WORKDIR /app

# Install Git (needed for Go modules)
RUN apk add --no-cache git

# Copy code and download dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binaries
RUN go build -o /bin/api ./cmd/api
RUN go build -o /bin/worker ./cmd/worker

# Default to api binary (can be overridden by docker-compose)
ENTRYPOINT ["/bin/api"]
