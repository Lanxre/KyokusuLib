FROM oven/bun:1 AS deps
WORKDIR /app
COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

FROM oven/bun:1 AS build
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
ARG NUXT_PUBLIC_API_BASE
ARG NUXT_API_INTERNAL_URL
ENV NUXT_PUBLIC_API_BASE=${NUXT_PUBLIC_API_BASE}
ENV NUXT_API_INTERNAL_URL=${NUXT_API_INTERNAL_URL}
RUN bun run build

FROM oven/bun:1 AS runner
WORKDIR /app
COPY --from=build /app/.output ./.output
EXPOSE 3000
CMD ["bun", ".output/server/index.mjs"]
