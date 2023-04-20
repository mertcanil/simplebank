init:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: init migrateup migratedown sqlc