ARG GOLANG_VERSION=1.24
ARG ALPINE_VERSION=3.21

# PREPARE
FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS prepare
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd/
COPY internal/ ./internal/

# BUILD
FROM prepare AS build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-s -w" -o ./bin/transactions-api ./cmd/api/main.go

# RELEASE
FROM alpine:${ALPINE_VERSION} AS release
WORKDIR /app
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
COPY ./.env ./.env
USER nonroot
COPY --from=build --chown=root:root --chmod=755 /app/bin/transactions-api /app/transactions-api
EXPOSE 8080
CMD ["/app/transactions-api"]

# MIGRATE
FROM prepare AS migrate
WORKDIR /app/migrations
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.24.3
COPY scripts/db/migrations/mysql .
CMD ["sh", "-c", "goose mysql \"$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME\" up"]
