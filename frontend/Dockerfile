# Build stage
FROM oven/bun:latest AS builder

WORKDIR /app

# Определяем аргументы сборки
ARG NUXT_PUBLIC_API_BASE
ARG NUXT_API_INTERNAL_URL
ENV NUXT_PUBLIC_API_BASE=$NUXT_PUBLIC_API_BASE
ENV NUXT_API_INTERNAL_URL=$NUXT_API_INTERNAL_URL

# Copy package files
COPY package.json bun.lockb* ./

# Install dependencies
RUN bun install

# Copy source files
COPY . .

# Build the Nuxt app
RUN bun run build

# Final stage
FROM oven/bun:latest

WORKDIR /app

# Copy the built output from the builder stage
COPY --from=builder /app/.output ./.output

# Expose the port Nuxt runs on (default 3000)
EXPOSE 3000

# Start the application
CMD ["bun", ".output/server/index.mjs"]
