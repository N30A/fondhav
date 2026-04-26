# Migrate

Install migrate:
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

Create a new migration file:
```bash
migrate create -dir migrations -ext sql <name_of_migration>
```

Run up migrations:
```bash
migrate -source file://migrations -database "postgres://user:password@localhost:5432/database?sslmode=disable" up
```

Run down migrations:
```bash
migrate -source file://migrations -database "postgres://user:password@localhost:5432/database?sslmode=disable" down
```
