FROM golang:1.24.5-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/main cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /build/main /app/main
COPY --from=builder /build/migrations /app/migrations

ENTRYPOINT ["/app/main"]