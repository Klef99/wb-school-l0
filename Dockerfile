FROM golang:1.24.5-alpine AS builder

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/main cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /buil/main /app/main

ENTRYPOINT ["/app/main"]