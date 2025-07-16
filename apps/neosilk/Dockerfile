# Build stage
FROM golang:latest AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src/apps/neosilk
RUN go build -o /bin/main ./cmd/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /bin/main .
ENTRYPOINT ["./main"]

