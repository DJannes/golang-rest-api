sqlc-gen:
	docker run --rm -v "%cd%:/src" -w /src/config kjconroy/sqlc generate

docker-up:
	docker compose -f devops/docker-compose.yaml up

docker-down:
	docker compose -f devops/docker-compose.yaml down
