# Build stage
FROM golang:1.23.4-alpine3.19 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o /go/bin/main .

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /go/bin .
COPY config ./config
COPY db/migrations ./db/migrations
COPY start.sh .
COPY wait-for.sh .

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]