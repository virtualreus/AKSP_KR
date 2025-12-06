# AKSP Health Monitoring App

Курсовая работа: архитектура клиент-серверного приложения для мониторинга здоровья и ведения медицинских записей.

## Архитектура

- API Gateway (Go + Chi)
- Auth Service (Go + Chi + PostgreSQL)
- Medical Service (Go + Chi + PostgreSQL)
- Frontend (Vue.js + TypeScript)

## Минимальный стек для MVP

- Go 1.21+, Chi
- Vue 3 + TypeScript
- PostgreSQL
- Docker, Docker Compose

## Структура репозитория

- `api-gateway/` - API Gateway сервис
- `auth-service/` - Сервис аутентификации
- `medical-service/` - Сервис медицинских данных
- `frontend/` - Vue.js фронтенд приложение
- `docs/` - Документация и диаграммы архитектуры
- `docker-compose.yml` - Конфигурация для запуска всего стека

## Запуск

```bash
# Запуск всех сервисов
docker-compose up -d

# Просмотр логов
docker-compose logs -f

# Остановка
docker-compose down
```

После запуска:

- Frontend: http://localhost:5173
- API Gateway: http://localhost:8080
- Auth Service: http://localhost:8081 (внутренний)
- Medical Service: http://localhost:8082 (внутренний)

## Документация

Полная документация находится в папке [`docs/`](./docs/):

- **[Диаграммы архитектуры](./docs/diagrams/)** - Все обязательные диаграммы для пояснительной записки:
  - Схема основных компонентов системы
  - Схема взаимодействия между компонентами
  - API и контракты между компонентами
  - ER-диаграмма базы данных
  - Схема физической инфраструктуры (Docker, CI/CD, Kubernetes)

Все диаграммы созданы в формате Mermaid и могут быть экспортированы в PNG для вставки в пояснительную записку.

## Функциональность

- ✅ Регистрация и авторизация пользователей (JWT)
- ✅ CRUD операции для медицинских записей
- ✅ CRUD операции для показателей здоровья
- ✅ Изоляция данных по пользователям
- ✅ Docker Compose для развертывания
