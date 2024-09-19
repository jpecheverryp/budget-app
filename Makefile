dev:lint
	templ generate
	go run ./cmd/web -port=$$BUDGET_PORT -db_url=$$BUDGET_DB_URL -db_auth_token=$$BUDGET_DB_AUTH_TOKEN

lint:
	templ fmt .
	go fmt ./...

