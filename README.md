## How to start

Normal

```
docker compose up -d
go run migrate/main.go
GO_ENV=dev go run main.go
```

Enable Hot Reload

```
go install github.com/cosmtrek/air@latest
GO_ENV=dev air
```
