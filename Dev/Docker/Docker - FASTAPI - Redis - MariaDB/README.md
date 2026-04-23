# FASTAPI - MARIADB - REDIS

## 1. DOCKER COMPOSE - MARIADB

```bash
docker compose -p "mariadb" -f docker-compose-mariadb.yml up --build
```

## 2. DOCKER COMPOSE - REDIS

```bash
docker compose -p "redis" -f docker-compose-redis.yml up --build
```

## 3. DOCKER COMPOSE - FASTAPI

```bash
docker compose -p "apiserver" -f docker-compose-apiserver.yml up --build
```
