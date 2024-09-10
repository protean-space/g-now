# g-now

## 環境構築

### PostgreSQLコンテナの起動

```sh
cd backend
docker compose up
```

### マイグレーション実行

```sh
cd backend
GO_ENV=dev go run migrate/migrate.go
```

### Seed実行

```sh
GO_ENV=dev go run db/seed/seed.go
```

### サーバー起動

```sh
cd backend
GO_ENV=dev go run main.go
```

http://localhost:8080/category
