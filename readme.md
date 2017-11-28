# to-go api

*Up and running*

`cd $GOPATH/../src`

clone this repo

`cd to-go`

`go get`

On Mac, `export GOBIN=$GOPATH/bin` sometimes comes in handy

edit config/database.yml with your mysql credentials

`./db_ setup` to create a database

`./db_ migrate` to create tables

`go build server.go`

`./server` start the server

the setup and migrate scripts along with the server binary can all be run with an environment as defined in config/database.yml, e.g. `./db_ setup production` or `./server staging` - running without an environment will default to development

install and run the client (coming soon)

*Resources*

[https://thenewstack.io/make-a-restful-json-api-go/](https://thenewstack.io/make-a-restful-json-api-go/)

[http://jinzhu.me/gorm/](http://jinzhu.me/gorm/)

