dev:
	docker compose -f docker-compose.dev.yaml up

run:
	CONFIG_FILE=./config.dev.yaml go run main.go
