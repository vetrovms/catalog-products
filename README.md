# API Каталог товарів  
Веб застосунок з API за протоколом HTTP для обслуговування списку товарів вітрини магазину.

Документація [swagger](https://app.swaggerhub.com/apis/VETROVMS_1/catalog_products_open_api_3_0/1.0.11)

UML згенеровано з yml через [скрипт](https://github.com/nlohmann/swagger_to_uml) та https://plantuml.com/starting
![alt text](docs/uml.png "UML")

## Налаштування та запуск

1. Експортувати змінні оточення

```bash
export  POSTGRES_PRODUCTS_API_PASSWORD=postgres && \
export  POSTGRES_PRODUCTS_API_USER=postgres && \
export  POSTGRES_PRODUCTS_API_DATABASE=catalog_products && \
export  POSTGRES_PRODUCTS_API_DSN="user=postgres password=postgres host=postgres_products_api port=5432 dbname=catalog_products sslmode=disable" && \
export  POSTGRES_PRODUCTS_API_URL="postgres://postgres:postgres@postgres_products_api:5432/catalog_products?sslmode=disable" && \
export  POSTGRES_DB=catalog_products && \
export  POSTGRES_USER=postgres && \
export  POSTGRES_PRODUCTS_API_PORTS=8080:8080 && \
export  POSTGRES_PRODUCTS_API_PORT=8080
```
2. З директорії `config` запустити докер `docker compose up`. Для локального запуску (без докера) потрібно змінити в змінних оточення хост з імені сервіса на localhost або 127.0.0.1 

> Запуск міграцій відбувається автоматично при запуску контейнера. Для запуску міграцій вручну (при релізі) виконати в контейнері команду
>  `migrate -database "$POSTGRES_PRODUCTS_API_URL" -path internal/database/migrations up`

Unit тести `go test tests/*`