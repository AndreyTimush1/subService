# 📦 Subscriptions Service (Go 1.25 + PostgreSQL)

REST-сервис для управления пользовательскими подписками.  
Реализован на **Go 1.25**, с хранением данных в **PostgreSQL**, сборка и запуск через **Docker Compose**.

---

## 🚀 Функциональность

- CRUD-операции над подписками:
  - `POST /subscriptions` — создать подписку  
  - `GET /subscriptions/:id` — получить по ID  
  - `PUT /subscriptions/:id` — обновить  
  - `DELETE /subscriptions/:id` — удалить  

- Подсчёт общей стоимости подписок:
  - `GET /subscriptions/total?user_id=&service_name=`  

- Логирование запросов (Gin)  
- Конфигурация через `.env`  
- Возможность расширить API документацией Swagger  

---

## ⚙️ Запуск

1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/yourname/subscriptions-service.git
   cd subscriptions-service
   ```

2. Поднять сервис и базу:
   ```bash
   docker-compose up --build
   ```

3. Сервис будет доступен по адресу:
   ```
   http://localhost:8080
   ```

---

## 🗄️ Миграции

Чтобы создать таблицу вручную, зайдите в контейнер PostgreSQL:

```bash
docker exec -it subscriptions-service-go125-postgres-1 psql -U postgres -d subscriptions
```

и выполните:

```sql
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE
);
```

(в будущем можно подключить **golang-migrate** для автоматического запуска).

---

## 📌 Примеры запросов

### Создание подписки
```bash
curl -X POST http://localhost:8080/subscriptions   -H "Content-Type: application/json"   -d '{
    "service_name": "Yandex Plus",
    "price": 400,
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
    "start_date": "2025-07-01T00:00:00Z"
  }'
```

### Получение по ID
```bash
curl http://localhost:8080/subscriptions/<uuid>
```

### Подсчёт общей суммы
```bash
curl "http://localhost:8080/subscriptions/total?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&service_name=Yandex Plus"
```

---

## 📂 Структура проекта

```
.
├── cmd/
│   └── main.go               # входная точка приложения
├── internal/
│   ├── db/                   # подключение к БД и миграции
│   ├── handlers/             # HTTP-обработчики (Gin)
│   ├── models/               # модели данных
│   └── repository/           # слой работы с БД
├── docker-compose.yml
├── Dockerfile
├── .env
└── README.md
```

---

## 🛠 Используемые технологии

- [Go 1.25](https://golang.org/)  
- [Gin](https://github.com/gin-gonic/gin)  
- [PostgreSQL](https://www.postgresql.org/)  
- [pgx](https://github.com/jackc/pgx)  
- [Docker Compose](https://docs.docker.com/compose/)  

---

⚡ Автор: *твой ник/имя*
