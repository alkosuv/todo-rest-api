# Todo 
Todo REST API server on golang

## Setting project
Создать в корне файл *.env* и добавить поля.

```
    BIND_ADDR = ""
    LOG_LEVEL = ""
    SESSION_KEY = ""

    DB_HOST = ""
    DB_PORT = ""
    DB_USER = ""
    DB_PASSWORD = ""
    DB_DATABASE = ""
    DB_SSLMODE = ""
```

## Tests

### Test api/v1/store/sqlstore
Для запуска тестов необходимо создать файл *.env* и добавить поля

```
DB_TEST_HOST = "192.168.0.20"
DB_TEST_PORT = "5432"
DB_TEST_USER = "postgres"
DB_TEST_PASSWORD = "0000"
DB_TEST_DATABASE = "todo_test"
DB_TEST_SSLMODE = "disable"
```
