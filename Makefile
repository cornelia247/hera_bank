startPostgres:
	docker start postgres12
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root hera_bank

dropdb:
	docker exec -it postgres12 dropdb hera_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/hera_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/hera_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...
.PHONY:  postgres createdb dropdb migrateup migratedown sqlc startPostgres

