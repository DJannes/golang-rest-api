sqlc-gen:
	docker run --rm -v "%cd%:/src" -w /src/config kjconroy/sqlc generate
