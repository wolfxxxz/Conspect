build:
	go build -v ./cmd/api/
create_migrations_repositories:
	migrate create -ext sql -dir migrations UsersCreationMigration
migration_up:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" up
migration_drop:
	migrate -path migrations -database "postgres://localhost:5433/words?sslmode=disable&user=postgres&password=1" down