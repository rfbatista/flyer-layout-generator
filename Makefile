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
docker-db:
	docker run --rm -d --name atlas-sqlc -p 5432:5432 -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=123 postgres
apply:
	atlas schema apply --url "postgres://admin:123@localhost:5432/algvisual?sslmode=disable" --dev-url "docker://postgres" --to "file://intelal/infrastructure/database/schema"
apply_in_server:
	atlas schema apply --url "postgres://postgres:123@localhost:5432/algvisual?sslmode=disable" --to "file://internal/infrastructure/database/schema"
clean:
	atlas schema clean --url "postgres://admin:123@localhost:5432/algvisual?sslmode=disable"
	PGPASSWORD=$(PGPASSWORD) psql -U admin -h localhost -p 5432 -d algvisual -c 'CREATE SCHEMA public;'
migrate:
	atlas migrate diff $(msg) \
		--dev-url "postgres://admin:123@localhost:5432/algvisual?search_path=public&sslmode=disable" \
		--dir "file://scripts/migrations" \
		--to "file://internal/infrastructure/database/schema"
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
buildai:
	docker build -t ai -f ./scripts/ai/Dockerfile.prod .
runaid:
	docker run -d -p 8080:8080 --network="host" --name ai ai
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

prod-ai-status:
	sudo systemctl status ai.service

prod-ai-restart:
	sudo systemctl restart ai.service

run_in_server:
	go run ./cmd/server/main.go
build_server:
	go build -o ./tmp/main ./cmd/server/main.go
restart:
	sudo systemctl restart server
status:
	sudo systemctl status server
stop:
	sudo systemctl stop server
restart_ai:
	sudo systemctl restart ai
status_ai:
	sudo systemctl status ai
log:
	journalctl -xe -u server.service

.PHONY: clean

