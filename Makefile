setup:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/http-swagger
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

run:
	swag init && go run main.go

sqlgen:
	docker run --rm -v "%cd%:/src" -w /src/config kjconroy/sqlc generate

up:
	docker compose -f devops/docker-compose.yaml up

down:
	docker compose -f devops/docker-compose.yaml down

dbup:
	migrate -verbose -source file://schema/migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable up

dbdown:
	migrate -verbose -source file://schema/migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable down --all

dbforce:
	migrate -verbose -source file://schema/migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable force $(version)
