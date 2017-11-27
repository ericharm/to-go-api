# to-go api

*Up and running*

clone this repo

`cd to-go`

`go get`

edit config/database.yml with your mysql credentials

`go run db_setup/db_setup.go` to create a database

`go run db_migrate/db_migrate.go` to create tables

`go build server.go`

`./server` start the server

the db_setup and db_migrate scripts, along with the server binary can all be run with an environment as defined in config/database.yml, e.g. `./server staging`

install and run the client (coming soon)

