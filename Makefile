run:
	go run ./cmd/main.go


postgres:
	docker run --name my_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

test:
	go test -v ./service/user


.PHONY: run postgres