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

# Быстрый старт
* Чтобы быстро развернуть апи, вместе с базой данных, можно воспользоваться docker-compose. Перед запуском docker-compose необходимо сделать конфигурирационный файл

Пример .env:
* Создайте .env файл в корне проекта
```dotenv
# ПРИМЕР
ENV=local
USER_DB=postgres
HOST_DB=test-task-db
PORT_DB=5432
PASSWORD_DB=1234
NAME_DB=TestTaskDB
DBMS=postgres
SSL_MODE=disable
API_PORT=8080
API_TIMEOUT=5s
EXTERNAL_API_URL_BASE=http://localhost:8080
```

> [!WARNING]
> Это важное предупреждение, которое нужно учитывать.
>
> Без файла .env апи не развернётся, будьте внимательны с заполнением его содержимого

* После создания конфигурационного файла, можно спокойно запустить docker-compose
```shell
make run-api-docker-compose-prod
```

# Как использовать?
* Создайте .env файл в корне проекта, или же в папке с бинарником АПИ
```dotenv
# prod local
# [OPTIONAL]
ENV=prod
USER_DB=user
HOST_DB=database
PORT_DB=5432
PASSWORD_DB=password
NAME_DB=dbname
DBMS=postgres
# enable disable
# [OPTIONAL]
SSL_MODE=enable
# [OPTIONAL]
API_PORT=8080
# [OPTIONAL]
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
