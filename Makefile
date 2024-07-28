postgres:
	docker run --name postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
createDb:
	docker exec -it postgres createdb --username=postgres simple_bank

dropDb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropDb migrateup migratedown sqlc test