migration-status:
	sql-migrate status
migrate:
	sql-migrate $0
models-dev:
	sqlboiler sqlite3
setup:
	go get -v github.com/rubenv/sql-migrate/...
	go install github.com/volatiletech/sqlboiler-sqlite3@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
