# Go boilerplate
This boilerplate follow Clean Architecture, work best for monolith Go application with multiple domain modules.


## Requirements

### Golangci-lint:
Installation: https://golangci-lint.run/usage/install/#local-installation

### Pre-commit hooks:
Installation: https://pre-commit.com/#installation

After install pre-commit, run the following commands:
```bash
# Run pre-commit install to install pre-commit into your git hooks. pre-commit will now run on every commit.
pre-commit install
pre-commit install --hook-type commit-msg
```

## How to create a DB Migration

### Prerequisite:
Install `goose` https://github.com/pressly/goose
```
go install github.com/pressly/goose/v3/cmd/goose@v3.7.0
```

### Create new file to migrate:
```bash
goose -dir ./db/migrations create add_order_table sql
```

```sql
-- +goose Up
-- +goose StatementBegin
CREATE TABLE "public"."order"
(
    "id"     varchar NOT NULL,
    "amount" varchar NOT NULL,
    CONSTRAINT "order_pk" PRIMARY KEY ("id")
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."order";
-- +goose StatementEnd
```

### Migrate: (Change dsn yourself)

```bash
goose -dir ./db/migrations postgres "postgresql://postgres:postgres@127.0.0.1:5435/db_name?sslmode=disable" up
```

# TODO:
[ ] CLI: run server, swagger file generate
# go-boilerplate
# go-boilerplate
