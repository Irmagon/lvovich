# fioincline — склонение русских ФИО и городов (Go)

Сервис склонения русских фамилий, имён, отчеств и названий городов по падежам.
Предоставляет REST, SOAP, WSDL и Swagger UI.

**Движок склонения — Go-порт библиотеки [nodkz/lvovich](https://github.com/nodkz/lvovich).**

---

## Благодарности / Acknowledgements

Правила склонения и определения пола, а также движок падежей заимствованы из
замечательной JavaScript-библиотеки **[nodkz/lvovich](https://github.com/nodkz/lvovich)**:

> Склонение русских ФИО и городов, выдержавшее проверку временем.
> Спасибо [@nodkz](https://github.com/nodkz) и всем контрибьюторам за огромную
> работу по сбору и проверке правил русского языка.

В свою очередь, `nodkz/lvovich` наследует правила из проекта
[Petrovich](https://github.com/petrovich/petrovich-js) (правила склонения ФИО).

Проект сохраняет MIT-лицензию с указанием авторства исходных правил и кода —
см. файл [LICENSE](./LICENSE).

---

## Возможности

- Склонение фамилии, имени и отчества по 6 падежам
- Склонение названий городов (предложный, родительный, винительный)
- Определение пола по ФИО (male / female / androgynous)
- REST + SOAP + WSDL + Swagger UI
- Конфигурация через `config.ini` (порт, токен, whitelist IP, Swagger)
- Асинхронный логгер запросов

## Быстрый старт

```
go build ./cmd/server
fioincline-server.exe
```

Или запустите собранный бинарник `fioincline-server.exe` в корне проекта.
Сервер запустится на адресе и порте из `config.ini` (по умолчанию `0.0.0.0:3000`).

## Endpoint-ы

| Интерфейс | Путь | Назначение |
|---|---|---|
| SOAP | `/soap` | SOAP-сервис |
| WSDL | `/wsdl` | WSDL-схема |
| REST | `/api/incline` | Склонение ФИО (JSON) |
| REST | `/api/gender` | Определение пола (JSON) |
| REST | `/api/city/in` | Город предложный (JSON) |
| REST | `/api/city/from` | Город родительный (JSON) |
| REST | `/api/city/to` | Город винительный (JSON) |
| Swagger | `/api-docs` | Интерактивная документация |

## Конфигурация

```ini
[server]
address = 0.0.0.0
port = 3000
swagger = true

[auth]
token = mysecret
allowed_ips = 127.0.0.1, ::1
```

Полное описание API и конфигурации — в файле [`service-doc.txt`](./service-doc.txt).

## Тесты

```
go test ./...
```

## Лицензия

MIT. См. файл [LICENSE](./LICENSE).