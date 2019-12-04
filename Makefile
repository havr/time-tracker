build:
	go build cmd/tracker/tracker.go

migrate-up:
	migrate -database postgres://${TT_DATABASE_URL} -source file://./migrations up

migrate-up:
	migrate -database postgres://${TT_DATABASE_URL} -source file://./migrations up 1

migrate-down:
	migrate -database postgres://${TT_DATABASE_URL} -source file://./migrations down 1

