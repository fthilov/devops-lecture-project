# Stage 1 - Build Stage
FROM golang:latest AS builder

ARG service

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main ./$service/cmd/main.go

# # Stage 2 - Final Stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/main/ .
EXPOSE 8080
CMD ["./main"]
