prequisite:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/http-swagger
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

sqlc-gen:
	docker run --rm -v "%cd%:/src" -w /src/config kjconroy/sqlc generate

docker-up:
	docker compose -f devops/docker-compose.yaml up

docker-down:
	docker compose -f devops/docker-compose.yaml down

migrate-up-x:
	migrate -verbose -source file://schema/migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable"&"search_path=master up

migrate-up:
	migrate -verbose -source file://schema/migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable up

migrate-down:
	migrate -verbose -source file://schema/migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable down --all

migrate-force:
	migrate -verbose -source file://schema/migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable force $(version)
