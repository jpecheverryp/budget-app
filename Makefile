dev:
	templ generate
	go run ./cmd/web -port=8080 -dsn=$$BUDGET_DB_DSN

lint:
	templ fmt .
	go fmt ./...
