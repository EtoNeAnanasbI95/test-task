# Описание
* Простенькая апи для тестового задания, в компанию Effective mobile

# Библиотеки
* gin
* zerolog
* godotenv
* cleanenv
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
* Если .env файл лежит в корне проекта или в одной директории с бинарником, то он подхватится автоматически, но можно явно указать путь до него с помощью флага `--config`
* Пример:
```shell
go run ./cmd/test-task/main.go --config ./.env
```
или

```shell
./test-task --config ./.env
```
