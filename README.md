# KyokusuLib

Платформа для чтения ранобэ и манги. SSR-фронтенд на Nuxt (Bun), REST API на Go, PostgreSQL.

---

## Стек

| Слой | Технологии |
|------|------------|
| Фронтенд | Nuxt 4 (Vue 3, SSR), Bun, Tailwind CSS 4, Pinia |
| Бэкенд | Go, gorilla/mux, godotenv |
| База | PostgreSQL 15 |
| Инфра | Docker Compose |

---

## Разработка

### Быстрый старт

```bash
cp .env.example .env
docker compose up --build -d
```

Фронтенд: `http://localhost:3000`
API: `http://localhost:8080`

### Локально (без Docker)

Только база в Docker, код с hot-reload:

```bash
docker compose up db -d

cd backend
go mod download
go install github.com/air-verse/air@latest
air

cd frontend
bun install
bun run dev
```

### Переменные окружения

Проект использует два `.env`:

- `backend/.env` — для локального запуска бэкенда (читается через godotenv)
- `.env` (корень) — для Docker Compose (подхватывается автоматически)

Оба содержат одно и то же. Основные переменные:

| Переменная | Описание |
|-----------|----------|
| `DATABASE_URL` | Подключение к PostgreSQL |
| `JWT_SECRET` | Секрет для JWT-токенов |
| `GOOGLE_CLIENT_ID` | OAuth Google Client ID |
| `GOOGLE_CLIENT_SECRET` | OAuth Google Client Secret |
| `DISCORD_CLIENT_ID` | OAuth Discord Client ID |
| `DISCORD_CLIENT_SECRET` | OAuth Discord Client Secret |
| `KYOKUSU_EMAIL_NAME` | SMTP-логин для отправки писем |
| `KYOKUSU_EMAIL_PASSWORD` | SMTP-пароль |
| `FRONTEND_URL` | URL фронтенда (для CORS) |
| `BACKEND_URL` | URL бэкенда (для редиректов) |
| `NUXT_PUBLIC_API_BASE` | Базовый URL API (для браузера) |

---

## Production

### 1. Запуск

```bash
git clone <repo> && cd KyokusuLib
cp .env.example .env

docker compose up --build -d
```

### 2. Reverse Proxy (Caddy)

В корне проекта есть готовый `Caddyfile`:

```
kyokusulib.com {
    @backend path /api/* /uploads/*
    handle @backend {
        reverse_proxy localhost:8080
    }
    handle {
        reverse_proxy localhost:3000
    }
    header {
        X-Content-Type-Options "nosniff"
        X-Frame-Options "DENY"
        X-XSS-Protection "1; mode=block"
        Referrer-Policy "strict-origin-when-cross-origin"
    }
}
```

```bash
sudo apt install caddy
sudo cp Caddyfile /etc/caddy/Caddyfile   # заменить домен
sudo systemctl restart caddy
```

Caddy автоматически получит SSL-сертификат через Let's Encrypt.

### 3. Продление OAuth callback'ов

В Google Cloud Console и Discord Developer Portal указать URL:

```
https://<домен>/api/auth/google/callback
https://<домен>/api/auth/discord/callback
https://<домен>/api/socials/google/callback
https://<домен>/api/socials/discord/callback
```

### 4. Авто-деплой (GitHub Actions)

```yaml
# .github/workflows/deploy.yml
name: Deploy
on:
  push:
    branches: [main]
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Deploy to VPS
        uses: appleboy/ssh-action@v1
        with:
          host: ${{ secrets.VPS_HOST }}
          username: ${{ secrets.VPS_USER }}
          key: ${{ secrets.VPS_SSH_KEY }}
          script: |
            cd KyokusuLib
            git pull
            docker compose up --build -d
```

---

## Структура проекта

```
├── .env.example           # Шаблон переменных окружения
├── Caddyfile              # Production reverse proxy config
├── docker-compose.yml     # Docker Compose (все сервисы)
│
├── frontend/
│   ├── app/
│   │   ├── assets/        # Стили, изображения
│   │   ├── components/    # Vue-компоненты
│   │   ├── composables/   # API-клиенты, бизнес-логика
│   │   ├── constants/     # Константы, утилиты
│   │   ├── layouts/       # Шаблоны страниц
│   │   ├── middleware/    # Middleware (auth, router guards)
│   │   ├── pages/         # Страницы (роутинг Nuxt)
│   │   ├── stores/        # Pinia (состояние: auth, settings)
│   │   ├── types/         # TypeScript-типы
│   │   └── utils/         # Вспомогательные функции
│   ├── public/            # Статические файлы
│   ├── Dockerfile         # Production-сборка
│   └── nuxt.config.ts     # Конфигурация Nuxt
│
├── backend/
│   ├── cmd/api/main.go    # Точка входа
│   ├── internal/
│   │   ├── config/        # Чтение конфигов
│   │   ├── database/      # Миграции, подключение к БД
│   │   ├── handlers/      # HTTP-обработчики
│   │   ├── middleware/     # CORS, JWT, аутентификация
│   │   ├── models/        # Структуры данных
│   │   ├── repositories/  # Работа с БД
│   │   └── services/      # Бизнес-логика
│   ├── migrations/        # SQL-миграции
│   ├── uploads/           # Загруженные файлы
│   ├── Dockerfile         # Multi-stage сборка
│   ├── go.mod
│   └── go.sum
```

---

## API

REST API, базовый URL: `http://localhost:8080/api`

| Раздел | Описание |
|--------|----------|
| `/api/auth/` | Регистрация, логин, OAuth (Google/Discord) |
| `/api/socials/` | Привязка социальных сетей к аккаунту |
| `/api/users/` | Профили, настройки, уровень, активность |
| `/api/novelas/` | Каталог новелл, тома, главы, контент |
| `/api/authors/` | Авторы |
| `/api/teams/` | Издательские команды, участники, заявки |
| `/api/comments/` | Комментарии к новеллам, лайки, репорты |
| `/api/bookmarks/` | Закладки пользователя |
| `/api/notifications/` | Уведомления |
| `/api/admin/` | Админ-панель |
| `/uploads/` | Статические файлы (аватарки, баннеры, постеры) |

---

## Лицензия

MIT
