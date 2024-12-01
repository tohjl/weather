# backend/Dockerfile
FROM golang:1.23.3-alpine

# Set working directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod ./
RUN go mod tidy

# Copy backend code
COPY . .

# Build Go application
RUN go build -o weather-report-go .

# Expose backend API port
EXPOSE 8080

# Command to run the backend
CMD ["./weather-report-go"]
