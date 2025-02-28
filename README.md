# Описание
* Простенькая апи для тестового задания, в компанию Effective mobile

# Библиотеки
* gin
* zerolog
* godotenv
* cleanenv
* migrate
* samber/slog-zerolog
* samber/slog-gin
* swaggo/gin-swagger

# Как использовать?
* Создайте .env файл в корне проекта, или же в папке с бинарником АПИ
```dotenv
ENV=prod
CONNECTION_STRING="postgres://user:password@localhost:5432/dbname?sslmode=disable"
API_PORT=8080
API_TIMEOUT=5s
EXTERNAL_API_URL_BASE=http://localhost
```
* Для полного запуска, есть прописанная цель в makefile, чтобы запустить апи вместе с миграциями, необходимо запустить api-run цель, и прокинуть две переменные `CONFIG_PATH` и `MIGRATIONS_PATH`. 

Пример:
```shell
make api-run CONFIG_PATH="./.env" MIGRATIONS_PATH="./migrations"
```
---
## Ручной запуск api
* Путь до .env файл нужно явно указать с помощью флага `--config`
* Пример:
```shell
go run ./cmd/test-task/main.go --config ./.env
```
или

```shell
./test-task --config ./.env
```

## Ручной запуск migrator
* Путь до .env файл нужно явно указать с помощью флага `--config`
* Путь до директории с миграциями нужно явно указать с помощью флага `--migrations`
* Пример:
```shell
go run ./cmd/migrate/main.go --config ./.env --migrations ./migrations
```
