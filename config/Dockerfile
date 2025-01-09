FROM golang:latest

ARG ARG_POSTGRES_PRODUCTS_API_PASSWORD
ARG ARG_POSTGRES_PRODUCTS_API_USER
ARG ARG_POSTGRES_PRODUCTS_API_DATABASE
ARG ARG_POSTGRES_PRODUCTS_API_DSN
ARG ARG_POSTGRES_PRODUCTS_API_URL
ARG ARG_POSTGRES_PRODUCTS_API_PORT
ARG ARG_LOG_PATH

ENV POSTGRES_PRODUCTS_API_PASSWORD $ARG_POSTGRES_PRODUCTS_API_PASSWORD
ENV POSTGRES_PRODUCTS_API_USER $ARG_POSTGRES_PRODUCTS_API_USER
ENV POSTGRES_PRODUCTS_API_DATABASE $ARG_POSTGRES_PRODUCTS_API_DATABASE
ENV POSTGRES_PRODUCTS_API_DSN $ARG_POSTGRES_PRODUCTS_API_DSN
ENV POSTGRES_PRODUCTS_API_URL $ARG_POSTGRES_PRODUCTS_API_URL
ENV POSTGRES_PRODUCTS_API_PORT $ARG_POSTGRES_PRODUCTS_API_PORT
ENV LOG_PATH $ARG_LOG_PATH

# системні пакети
RUN apt-get update && \
    apt-get -y install --no-install-recommends --no-install-suggests \
    vim \
    nano \
    less \
    git \
    lsb-release \
    curl

# утиліта для міграцій golang-migrate
RUN wget http://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.deb && \
    dpkg -i migrate.linux-amd64.deb

# очистка
RUN apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# гіт
COPY git /git
RUN rm -rf /app \
    && mkdir -p /app \
    && cd /app \
    && git init \
    && cp -r /git /app/.git \
    && git remote -v \
    && git remote add origin https://github.com/vetrovms/catalog-products.git \
    && git remote -v \
    && git fetch origin master \
    && git checkout master \
    && rm -rf /git \
    && git config --global user.name "Thread" \
    && git config --global user.email thread@vipdesign.com

WORKDIR /app

# GO install
RUN go get -u github.com/gofiber/fiber/v3 && \
    go get gorm.io/gorm && \
    go get gorm.io/driver/postgres && \
    go get github.com/subosito/gotenv && \
    go mod tidy

# білд
RUN go build -C cmd -o app

# міграції та запуск застосунку
ENTRYPOINT migrate -database "$POSTGRES_PRODUCTS_API_URL" -path internal/database/migrations up && cmd/app