include .env

SHELL := /bin/bash
PGPASSWORD=123
VENV_DIR = venv
PYTHON = python
REQUIREMENTS = requirements.txt

run:
	PYTHONPATH=. uvicorn app.main:app --reload
upgrade:
	alembic upgrade head
sql:
	sqlc generate
db:
	docker run --rm -d --name atlas-sqlc -p 5432:5432 -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=123 postgres
apply:
	atlas schema apply --url "postgres://admin:123@localhost:5432/algvisual?sslmode=disable" --dev-url "docker://postgres" --to "file://internal/database/schema"
apply_in_server:
	atlas schema apply --url "postgres://admin:123@localhost:5432/postgres?sslmode=disable" --to "file://internal/database/schema"
clean:
	atlas schema clean --url "postgres://admin:123@localhost:5432/algvisual?sslmode=disable"
	PGPASSWORD=$(PGPASSWORD) psql -U admin -h localhost -p 5432 -d algvisual -c 'CREATE SCHEMA public;'
migrate:
	atlas migrate diff $(msg) \
		--dev-url "postgres://admin:123@localhost:5432/algvisual?search_path=public&sslmode=disable" \
		--dir "file://scripts/migrations" \
		--to "file://internal/database/schema"
server:
	go run ./cmd/server/main.go
dev-build:
	docker compose -f ./scripts/docker-compose.dev.yaml up --build
dev:
	docker compose -f ./scripts/docker-compose.dev.yaml up ai-dev postgres-dev
devlog:
	docker compose -f ./scripts/docker-compose.logs.yaml up
test:
	go test ./internal/...
usecase:
	hygen usecase new
down:
	docker compose -f ./scripts/docker-compose.dev.yaml down
ssh:
	ssh -i ./ssh-key.pem ec2-user@54.221.241.11
run_in_server:
	/usr/local/go/bin/go run ./cmd/server/main.go
build_in_server:
	/usr/local/go/bin/go build -o ./server ./cmd/server/main.go
restart:
	sudo systemctl restart server
status:
	sudo systemctl status server
buildai:
	docker build -t ai -f ./scripts/ai/Dockerfile.prod .
runaid:
	docker run -d -p 8080:8080 -v /home/ec2-user/alg_visual:/home/ec2-user/alg_visual -v /home/ec2-user:/home/ec2-user --network="host" --name ai ai
runai:
	uvicorn app.main:app --host 0.0.0.0 --port 8080 --reload
air:
	air
algai:
	docker build -t algvisual-ai:latest -f ./scripts/ai/Dockerfile.dev .
run-algai:
	docker run -v .:/app --net=host -p 8080:8080 -it algvisual-ai:latest
ai-poetry:
	python -m poetry run uvicorn app.main:app --host 0.0.0.0 --port 8080 --reload
ai:
	uvicorn app.main:app --host 0.0.0.0 --port 8080 --reload
db:
	docker compose up --build
activate:
	@echo "Activating virtual environment"
	@source $(VENV_DIR)/bin/activate
environment:
	shell python -m venv venv

.PHONY: clean
