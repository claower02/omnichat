# OmniChat CRM Integration

Мультиканальная платформа для интеграции любых мессенджеров (Telegram, WhatsApp) и любой CRM системы.

## Архитектура
Проект использует монорепозиторий:
- **`backend/`**: Go (Gin, WebSockets, GORM). Обрабатывает Webhook-и от мессенджеров, отдает API для UI и отправляет Webhook-и в вашу CRM.
- **`frontend/`**: React (Vite, TypeScript, TailwindCSS, Zustand). Премиальный UI для операторов с поддержкой боковой панели карточки клиента из CRM.

## Локальный запуск
Убедитесь, что у вас установлен Docker и Docker Compose.

```bash
docker-compose up --build
```
- Frontend доступен по адресу: http://localhost:3000
- Backend API доступен по адресу: http://localhost:8080
- База данных: PostgreSQL на порту 5432

## Развертывание на Railway.app
Проект готов к деплою на Railway. Поскольку это монорепозиторий, вам нужно:
1. Создать новый проект в Railway (Deploy from GitHub repo).
2. Выбрать репозиторий с этим кодом.
3. Railway найдет оба Dockerfile. Если нет, создайте 2 сервиса (Empty Service):
   - **Frontend Service**:
     - В настройках сервиса укажите `Root Directory` -> `/frontend`
     - Railway автоматически соберет Dockerfile и запустит Nginx.
   - **Backend Service**:
     - В настройках сервиса укажите `Root Directory` -> `/backend`
     - Railway соберет Go приложение.
4. Добавьте **PostgreSQL** базу данных через дашборд Railway.
5. Свяжите базу данных с Backend Service (передайте переменную `DATABASE_URL`).
6. Готово!

## Интеграция с любой CRM
1. Ваша CRM должна отправлять POST запросы на `/crm/v1/send`, когда менеджер пишет клиенту.
2. Backend будет отправлять POST Webhook с событиями (новое сообщение, смена статуса) на Webhook URL вашей CRM (настраивается).
