# GIN APP - MARIADB - REDIS

## 1. KHỞI TẠO GO MODULE TRONG THƯ MỤC SOURCE CODE

```bash
go mod init vmk-gin-app-docker &&
go get github.com/gin-gonic/gin &&
go get github.com/gin-contrib/cors &&
go get github.com/redis/go-redis/v9 &&
go get github.com/go-sql-driver/mysql
```

## 2. DOCKER COMPOSE - MARIADB

```bash
docker compose -p "mariadb" -f docker-compose-mariadb.yml up --build
```

## 3. DOCKER COMPOSE - REDIS

```bash
docker compose -p "redis" -f docker-compose-redis.yml up --build
```

## 4. DOCKER COMPOSE - GIN APP - DEVELOPMENT

```bash
docker compose -p "gin-app-dev" -f docker-compose-gin-app-dev.yml up --build --watch
```

## 5. DOCKER COMPOSE - GIN APP - PRODUCTION

```bash
docker compose -p "gin-app-prod" -f docker-compose-gin-app-prod.yml up --build
```
