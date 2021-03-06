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
go get ./...
make build
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

### Integration tests

I added some examples of DB integration tests how I see it in a very simple case. 
Of course, for production there should be much complex solution that speeds the process up.

```bash
export TT_TEST_DB=postgres://test:test@localhost:5432/time_tracker_test?sslmode=disable
docker run -d --name=tt_test_pg --tmpfs /var/lib/postgresql/data:rw -p 5432:5432 -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -e POSTGRES_DB=time_tracker_test postgres 
migrate -database ${TT_TEST_DB} -source file://./migrations up
go test -v ./...
```

### Environment variables

```
# required, postgres connection string
export TT_DATABASE_URL=postgres://postgres:verystr0ngpassword@localhost:5432/time_tracker?sslmode=disable 
# host to serve everything at
export TT_SERVE_AT=:8080
# optional, tells the app to serve the static dir
export TT_STATIC_DIR=ui/build 
# optional, specifies the directory with migrations to apply before the application starts
export TT_MIGRATE_FROM=./migrations 
# optional, wait the amount of seconds if the database returns errors on application start
export TT_WAIT_FOR_DATABASE_SECONDS=5 
```
