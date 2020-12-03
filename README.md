# Go Base Backend
### Built with Gin, Docker and Postgres

using cheeseburger as main branch to avoid using main

### Purpose
- Building a base backend so that we can fork it in the future to avoid building backend from scratch

### Features
- jwt authentication
- basic CRUD implementation
- dockerized
- docker compose with database service
- CI/CD (currently working in progress)

### Usage
#### development
- First, you need to run the postgres database
- `docker run --name postgresdb -e POSTGRES_PASSWORD=postgrespwd -e POSTGRES_USER=postgres -p 5432:5432 -d postgres`
- Then, install all necessary package if this is the first time you run
- `go get`
- Finally, run the program
- `go run main.go`

#### production
- Simply set up the variable you want in `docker-compose.yml`, then run
- `docker-compose up -d`

License
----
MIT
