run:
	PYTHONPATH=. uvicorn app.main:app --reload
upgrade:
	alembic upgrade head
sql:
	sqlc generate
db:
	docker run --rm -d --name atlas-sqlc -p 5432:5432 -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=123 postgres
apply:
	atlas schema apply --url "postgres://admin:123@localhost:5432/algvisual?sslmode=disable" --dev-url "docker://postgres" --to "file://database/schema"
clean:
	atlas schema clean --url "postgres://admin:123@localhost:5432/algvisual?sslmode=disable" 
migrate:
	atlas migrate diff $(msg) \
		--dev-url "postgres://admin:123@localhost:5432/algvisual?search_path=public&sslmode=disable" \
		--dir "file://database/migrations" \
		--to "file://database/schema"
server:
	go run ./cmd/server/main.go

dev-build:
	docker compose -f ./scripts/docker-compose.dev.yaml up --build
dev:
	docker compose -f ./scripts/docker-compose.dev.yaml up 
