# Todo 
Todo REST API server on golang

## URL

| URL                                             | Method | Description                                                       |
|-------------------------------------------------|--------|-------------------------------------------------------------------|
| /api/v1/sessions                                | post   | Создать сессию                                                    |
|                                                 | delete | Удалить сессию                                                    |
| /api/v1/users                                   | post   | Регистрация пользователя                                          |
| /api/v1/private/todos                           | get    | Получить все задачи пользователя                                  |
|                                                 | post   | Добавить задачу                                                   |
| /api/v1/private/todos/(todoID)                  | patch  | Обновить данные задачи                                            |
|                                                 | delete | Удалить задачу                                                    |
| /api/v1/private/todos/find?completed=false      | post   | Получить завершённые (false) или активные (true) задачи           |
| /api/v1/private/todos/count                     | get    | Получить количество записей в таблице                             |
| /api/v1/private/todos/find/count?completed=true | get    | Получить количество завершённых (false) или активных (true) задач |
| /api/v1/private/users                           | get    | Получить данные пользователя (авторизованного)                    |
|                                                 | patch  | Обновить данные пользователя                                      |
___

### Описание

* /api/v1/sessions - post

```JSON
    {
        "login": "user0",
        "password": "12345"
    }
```


* /api/v1/users - post

```JSON
    {
        "login": "user0",
        "password": "12345"
    }
```

* api/v1/private/todos - post

```JSON
    {
        "title": "Title"
    }
```

* api/v1/private/todos - patch

```JSON
{
    "column": "completed",
    "value": "true"
}
```

```JSON
{
    "column": "title",
    "value": "new title"
}
```

* api/v1/private/users - patch

```JSON
{
    "column": "name",
    "value": "name"
}
```

```JSON
{
    "column": "password",
    "value": "password"
}
```
___

## Makefile
*Команды работают на __linux и mac os__*

Для корректной работы __make__ необходимо заполнить поля в *Makefile* для доступа к БД
```
user = ""
password = ""
host = ""
port = ""
sslmode = ""
db_dev = ""
db_test = ""
```

+ Собрать и запустить проект
``` cmd
    $ make
```

+ Собрать и запустить проект (default)
``` cmd
    $ make start
```

+ Собрать проект
``` cmd
    $ make build
```

+ Запустить тесты
``` cmd
    $ make test
```

+ Создать миграцию
```cmd
    $ make migrate-create
```

+ Создать таблицы для баз  test и dev
``` cmd
    $ make migrate-up
```

 + Создать таблицы для dev
```cmd
    $ make migrate-dev-up
``` 

 + Создать таблицы для test
```cmd
    $ make migrate-test-up
 ``` 

 + Удалить таблицы в dev и test
```cmd
    $ make migrate-down
```

+ Удалить таблицу dev
```cmd
    $ make migrate-dev-down
```

+ Удалить таблицу test
```cmd
    $ make migrate-test-down
```

___

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

### Test
Для запуска тестов необходимо создать файл *.env* и добавить поля

```
DB_TEST_HOST = "192.168.0.20"
DB_TEST_PORT = "5432"
DB_TEST_USER = "postgres"
DB_TEST_PASSWORD = "0000"
DB_TEST_DATABASE = "todo_test"
DB_TEST_SSLMODE = "disable"
```
