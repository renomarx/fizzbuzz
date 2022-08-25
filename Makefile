
init-env:
	./scripts/init_env.sh

build:
	docker-compose build

dev: build init-env
	docker-compose up -d

build-base:
	docker build . -f build/Dockerfile --target builder -t fizzbuzz_base:latest

tests: build-base
	# docker-compose up -d postgres
	docker run -t --env REDIS_ADDR="127.0.0.1:6379" --network host fizzbuzz_base:latest go test -cover ./...

doc:
	swag init -g cmd/fizzbuzz/main.go
