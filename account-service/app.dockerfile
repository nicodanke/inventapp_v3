# Build stage
FROM golang:1.23.4-alpine3.19 AS builder
WORKDIR /app
COPY account-service account-service
COPY account-service/go.mod account-service/go.sum ./
RUN go build -o /go/bin/main ./account-service/main.go

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /go/bin .
COPY account-service/dev.env .
COPY account-service/db/migrations ./db/migrations
COPY start.sh .
COPY wait-for.sh .

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]