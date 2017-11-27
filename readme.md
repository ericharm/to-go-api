# to-go api

*Up and running*

clone this repo

`cd to-go`

`go get`

edit config/database.yml with your mysql credentials

`go run db_setup.go development` to create a database

`go run db_migrate.go development` to create tables

`go build main/init.go`

`./init` start the server

install and run the client (coming soon)

