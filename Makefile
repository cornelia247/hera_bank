startPostgres:
	docker start postgres12
postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root hera_bank

dropdb:
	docker exec -it postgres12 dropdb hera_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/hera_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/hera_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/hera_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/hera_bank?sslmode=disable" -verbose down 1

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/cornelia247/hera_bank/db/sqlc Store
	
.PHONY:  postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc startPostgres server mock

