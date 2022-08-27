.PHONY: init-env build dev build-base tests doc

init-env:
	./scripts/init_env.sh

build: init-env
	docker-compose build

dev: build
	docker-compose up -d

build-base:
	docker build . -f build/Dockerfile --target builder -t fizzbuzz_base:latest

tests: build-base
	docker run -t fizzbuzz_base:latest go test -cover ./...

doc:
	swag init -g cmd/fizzbuzz/main.go
