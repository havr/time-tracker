### What is it

It's a test assignment for Pento.

### How to run

```bash
docker build -t havr/time-tracker .
docker-compose up
```

It starts a server on a :8080 port.

# Migrations
I suggest to use the [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) tool.
The Makefile already contains the examples of its usage.

## How to build backend
```bash
go get
go build cmd/tracker/tracker.go
```

## Frontend
```bash
cd ui
yarn install
make build # just builds a web app
# or
make serve # starts a development server
```

The server serves by default on the 3333 port and proxies API calls to the backend it expects to run on the 8080.

### Environment variables

```
TT_DATABASE_URL=postgres://postgres:verystr0ngpassword@postgres:5432/time_tracker?sslmode=disable # required, postgres connection string
TT_SERVE_AT=:8080 # host to serve everything at
TT_STATIC_DIR=ui/build # optional, tells the app to serve the static dir
TT_MIGRATE_FROM=./migrations # optional, specifies the directory with migrations to apply before the application starts
TT_WAIT_FOR_DATABASE_SECONDS=5 # optional, wait the amount of seconds if the database returns errors on application start
```
