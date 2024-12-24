# API Каталог товарів

Веб застосунок з API за протоколом HTTP для обслуговування списку товарів вітрини магазину.

## Налаштування та запуск
1. Експортувати змінні оточення
> export POSTGRES_PRODUCTS_API_PASSWORD=postgres && \
export POSTGRES_PRODUCTS_API_USER=postgres && \
export POSTGRES_PRODUCTS_API_DATABASE=catalog_products && \
export POSTGRES_PRODUCTS_API_DSN="user=postgres password=postgres host=postgres_products_api port=5432 dbname=catalog_products sslmode=disable" && \
export POSTGRES_PRODUCTS_API_URL="postgres://postgres:postgres@postgres_products_api:5432/catalog_products?sslmode=disable" && \
export POSTGRES_DB=catalog_products && \
export POSTGRES_USER=postgres && \
export POSTGRES_PRODUCTS_API_PORTS=8080:8080
2. Локально з директорії `config` запустити докер `docker compose up`, або завантажити вміст директорії на сервер та запустити там докер.

> Запуск міграцій відбувається автоматично при запуску контейнера. Для запуску вручну (при релізі) виконати в контейнері команду
> `migrate  -database  "$POSTGRES_PRODUCTS_API_URL"  -path  internal/database/migrations  up`