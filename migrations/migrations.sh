goose -dir ./migrations postgres "postgres://user:password@localhost:5432/banks?sslmode=disable" status
goose -dir ./migrations postgres "postgres://user:password@localhost:5432/banks?sslmode=disable" up