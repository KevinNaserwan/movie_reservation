migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/db_movie_user?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/db_movie_user?sslmode=disable" -verbose down

sqlc:
	sqlc generate
