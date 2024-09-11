dev:
	templ generate
	go run ./cmd/web

lint:
	templ fmt .
	go fmt ./...
