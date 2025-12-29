FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o authd ./services/auth/cmd/authd

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/authd .

EXPOSE 50051

CMD ["./authd"]
