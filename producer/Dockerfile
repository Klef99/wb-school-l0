FROM golang:1.24.5-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/main .

FROM scratch

WORKDIR /app

COPY --from=builder /build/main /app/main

ENTRYPOINT ["/app/main"]