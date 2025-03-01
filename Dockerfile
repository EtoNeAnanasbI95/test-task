FROM golang:1.23-alpine AS BUILDER
LABEL authors="notkirilov"

WORKDIR /test-task-api

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o test-task-api ./cmd/test-task/main.go
RUN go build -o test-task-migrator ./cmd/migrate/main.go

FROM golang:1.23-alpine AS RUNNER
WORKDIR /test-task-api

COPY --from=BUILDER /test-task-api/test-task-api ./test-task-api
COPY --from=BUILDER /test-task-api/test-task-migrator ./test-task-migrator
COPY --from=BUILDER /test-task-api/migrations ./migrations
COPY --from=BUILDER /test-task-api/.env ./.env

EXPOSE 8080

CMD ["./test-task-migrator --config ./.env --migrations ./migrations && ./test-task-api --config ./.env"]
