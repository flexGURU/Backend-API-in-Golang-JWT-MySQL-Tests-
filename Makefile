run:
	go run ./cmd/main.go


postgres:
	docker run --name my_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

test:
	go test -v ./service/user

createdb:
	docker exec -it my_postgres createdb --username=root demogo

dropdb:
	docker exec -it my_postgres psql -U root -d demogo -c "\dt"

	 

accesspg:
	docker exec -it my_postgres psql -U root -d demogo
	


migrateup:
	migrate -path cmd/migrate/migrations -database "postgresql://root:secret@localhost:5432/demogo?sslmode=disable" -verbose up


migratedown:
	migrate -path cmd/migrate/migrations -database "postgresql://root:secret@localhost:5432/demogo?sslmode=disable"  -verbose down


migratecreate:
	migrate create -ext sql -dir cmd/migrate/migrations -seq add-user
	migrate create -ext sql -dir cmd/migrate/migrations -seq add-product
	migrate create -ext sql -dir cmd/migrate/migrations -seq add-orders
	migrate create -ext sql -dir cmd/migrate/migrations -seq add-order-items


.PHONY: run postgres migrateCreate createdb dropdb migrateup migratedown