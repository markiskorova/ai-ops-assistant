FROM golang:1.23

WORKDIR /app

# Copy the entire project
COPY . .

# Download dependencies
RUN go mod tidy

# Default command is API; overridden in docker-compose for worker
CMD ["go", "run", "cmd/api/main.go"]