## How to start

Normal

```
docker compose up -d
GO_ENV=dev go run migrate/migrate.go
GO_ENV=dev go run main.go
```

Enable Hot Reload

```
go install github.com/cosmtrek/air@latest
GO_ENV=dev air
```
