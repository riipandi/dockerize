# ------------------------------------------------------------------------------
# docker build -t laravel:latest -t laravel:1.0.0 .
# docker run -d --name laravel -p 8000:8000 laravel:latest
# docker-compose run --rm laravel php artisan migrate:fresh --seed
# ------------------------------------------------------------------------------

# ------------------------------------------------------------------------------
# Install NPM dependencies only when needed
# ------------------------------------------------------------------------------
FROM node:16-alpine AS npmdeps
WORKDIR /app
COPY . .
RUN npm install --frozen-lockfile
RUN npm run production

# ------------------------------------------------------------------------------
# Install NPM dependencies only when needed
# ------------------------------------------------------------------------------
FROM composer:latest AS builder
WORKDIR /app
COPY . .
COPY --from=npmdeps /app/package-lock.json ./package-lock.json
COPY --from=npmdeps /app/public ./public
RUN composer install

# ------------------------------------------------------------------------------
# Install NPM dependencies only when needed
# ------------------------------------------------------------------------------
FROM php:cli-alpine AS runner
LABEL maintainer="Aris Ripandi"
ENV NODE_ENV production
WORKDIR /app

COPY . .
COPY --from=npmdeps /app/package-lock.json ./package-lock.json
COPY --from=builder /app/composer.lock ./composer.lock
COPY --from=builder /app/vendor ./vendor
COPY --from=builder /app/public ./public

RUN php artisan key:generate --force
RUN php artisan storage:link

EXPOSE 8000

CMD ["php", "artisan", "serve", "--host", "0.0.0.0", "--port", "8000"]
