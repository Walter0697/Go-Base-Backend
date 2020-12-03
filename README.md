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

### Notes
- In `main.go`, we create a new user here to make sure that we can login, you can remove it after creating your first user
- `/allusers` this get request is only for checking if the user is being added or not, you can remove it as you wish
- all object will return all their fields that matches their database value, please wrap it to avoid any potential leaks

### Issue
- request body only accept raw json
- unmatching field will return error with object name instead of json field name


License
----
MIT
