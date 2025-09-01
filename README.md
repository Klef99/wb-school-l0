# Техношкола Wildberries: задание L0

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/klef99/wb-school-l0)
![Report](https://goreportcard.com/badge/github.com/klef99/wb-school-l0)

## Задание

Данное задание предполагает создание небольшого микросервиса на Go с использованием базы данных и очереди сообщений.
Сервис будет получать данные заказов из очереди (Kafka), сохранять их в базу данных (PostgreSQL) и кэшировать в памяти
для быстрого доступа.

Данные, приходящие из очереди, могут быть невалидными — необходимо предусмотреть обработку ошибок (например, игнорируйте
или логируйте некорректные сообщения). В ходе реализации убедитесь, что при сбоях (ошибка базы, падение сервиса) данные
не теряются — используйте транзакции, механизм подтверждения сообщений от брокера и т.д.

## Формат данных

```json
{
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
```

## Локальный запуск

Клонировать проект

```bash
  git clone http://github.com/klef99/wb-school-l0
```

Перейти в директорию сервиса

```bash
  cd wb-school-l0
```

Установить зависимости

```bash
  make deps-bin
  make dotenv
  make gen-wire
```

Запустить сервис и инфраструктуру

```bash
  make up
```

Данная команда запустит docker-compose окружение со следующими сервисами:

    1. Backend
    2. Frontend
    3. Kafka (3 реплики)
    4. PostgreSQL
    5. Redis

Доступны следующие адреса:

1. `localhost:8080` - frontend
2. `localhost:8085` - backend

Запуск приложения вне docker:

```bash
make run
```

Запуск приложения вне docker требует настроенной инфраструктуры, записанных в .env данных и настроенной конфигурации
nginx reverse proxy (config/nginx/nginx.conf)

## Доступные команды

```bash
Usage:
  make <target>

Targets:
  deps-bin            Install service dependencies and tools
  lint                Run linter
  dotenv              Generate .env file from .env.example
  gen-wire            Generate wire_gen.go file
  migrations          Run migrations utility
  up                  Up service and all necessary infrastructure
  down                Down service and all necessary infrastructure
  rm                  Remove service and all necessary infrastructure
  run                 Build and run application (go run) (infrastructure should exits)
```

## Backend API

#### Health

```http
  GET /v1/health
```

**Ответы**

| Код   | Описание                 |
|:------|:-------------------------|
| `200` | Сервис готов к работе    |
| `500` | Сервис не готов к работе |

#### Get order

```http
  GET /v1/orders/${id}
```

**Параметры запроса**

| Параметр | Тип      | Описание                  |
|:---------|:---------|:--------------------------|
| `id`     | `string` | **Требуется**. UID заказа |

**Ответы**

| Код   | Описание                  |
|:------|:--------------------------|
| `200` | Успешно                   |
| `404` | Заказ не был найден       |
| `500` | Внутренняя ошибка сервиса |

## Технологии

**Backend:** Go, Kafka, Redis, Echo, Wire, Nginx