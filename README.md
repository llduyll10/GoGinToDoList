# GoGinToDoList
To do list application using Gin framework


## How To Use
## Prerequisite 🏆
- Go Version `>= go 1.20`
- PostgreSQL Version `>= version 15.0`

### With Docker
1. Copy the example environment file and configure it:
  ```bash
  cp.env.example .env
  ```
2. Build Docker
  ```bash
  docker-compose build --no-cache
  ```
3. Run Docker Compose
  ```bash
  docker-compose up -d
  ```
## Run Migrations and Seeder
To run migrations and seed the database, use the following commands:

```bash
go run main.go --migrate --seed
```

#### Migrate Database
To migrate the database schema
```bash
go run main.go --migrate
```
This command will apply all pending migrations to your PostgreSQL database specified in `.env`

#### Seeder Database
To seed the database with initial data:
```bash
go run main.go --seed
```
This command will populate the database with initial data using the seeders defined in your application.