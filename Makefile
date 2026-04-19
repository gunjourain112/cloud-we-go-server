.PHONY: run-gin run-hertz swag-gin ent-gin docker-up docker-down

run-gin:
	cd gin && go run main.go

ent-gin:
	cd gin && go generate ./...

swag-gin:
	cd gin && swag init -g main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down
