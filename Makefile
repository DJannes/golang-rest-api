go-prequisite:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

swag-prequisite:
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/http-swagger

sqlc-gen:
	docker run --rm -v "%cd%:/src" -w /src/config kjconroy/sqlc generate

docker-up:
	docker compose -f devops/docker-compose.yaml up

docker-down:
	docker compose -f devops/docker-compose.yaml down
