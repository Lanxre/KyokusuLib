# 📚 KyokusuLib

**KyokusuLib** — это высокопроизводительная платформа для чтения ранобэ и манги. Проект использует современный стек: **Nuxt** на рантайме **Bun** для фронтенда и **Golang** с роутером **gorilla/mux** для бэкенда.

## ⚡ Технологический стек

### Frontend
- **Framework:** [Nuxt](https://nuxt.com/) (Vue 3, SSR)
- **Runtime:** [Bun](https://bun.sh/)
- **Styling:** [Tailwind CSS 4](https://tailwindcss.com/)
- **State:** Pinia & VueUse

### Backend
- **Language:** [Golang](https://go.dev/)
- **Router:** [gorilla/mux](https://github.com/gorilla/mux)
- **Architecture:** RESTful API

---

## 🚀 Особенности

- **Full SSR:** Быстрая первая отрисовка (LCP) и SEO-оптимизация.
- **Ultra-fast Bundle:** Использование Bun для сборки и Tailwind 4 через Vite-плагин.
- **Smart Color Mode:** Автоматическое определение темы на сервере (через Cookies) и мгновенное применение без «вспышек» белого.
- **User Progression:** Система уровней, опыта (XP) и кастомных тегов профиля.
- **Complex Content:** Поддержка многоуровневой структуры (Новелла -> Тома -> Главы).

---

## 📁 Структура репозитория

```bash
├── frontend/       # Nuxt Frontend
│   ├── components/       # Атомарные Vue компоненты
│   ├── composables/      # Бизнес-логика и API-клиенты
│   └── stores/           # Pinia (auth, activity, settings)
├── backend/       # Golang Backend
│   ├── cmd/              # Точки входа (main.go)
│   ├── pkg/              # Роуты, модели, хендлеры
│   └── internal/         # Бизнес-логика и БД
```

---

## 🛠 Установка и запуск

### Фронтенд (Bun)

1. Установите зависимости:
   ```bash
   bun install
   ```
2. Настройте `.env`:
   ```env
   NUXT_PUBLIC_API_BASE=http://localhost:8080
   ```
3. Запустите в режиме разработки:
   ```bash
   bun run dev
   ```

### Бэкенд (Golang)

1. Перейдите в папку API:
   ```bash
   cd backend
   ```
2. Скачайте зависимости:
   ```bash
   go mod download
   ```
3. Запустите сервер:
   ```bash
   go run cmd/main.go
   ```

---

## 🔧 Конфигурация Proxy

Для обхода CORS в разработке Nuxt настроен на проксирование запросов к Go-серверу:

```typescript
routeRules: {
  "/api/**": { proxy: "http://localhost:8080/api/**" }
}
```

---

## 🎨 Дизайн

- **Темы:** Поддержка системной, светлой и темной тем.
- **Эффекты:** Радиальные градиенты Tailwind 4 для темной темы (`dark:bg-radial-[at_center]`).
- **Адаптивность:** Mobile-first подход, адаптированные карусели и сетки контента.

## 🔒 Безопасность

- Проверка прав доступа на фронтенде (`hasPermission`).
- Middleware на стороне Go для валидации JWT/Session.
- Защита от гидратации (Hydration mismatch) через `ClientOnly` и `useAsyncData`.
