# to-go api

### Up and running

`cd $GOPATH/src`

`git clone https://github.com/ericharm/to-go.git`

`cd to-go`

edit config/database.yml with your mysql credentials

`go get`

`go build`

On Mac, `export GOBIN=$GOPATH/bin` sometimes comes in handy

`./db_ setup` to create a database

`./db_ migrate` to create tables

`./db_ seed`

`./to-go` start the server

the db&#95; scripts along with the server binary can all be run with an environment as defined in config/database.yml, e.g. `./db_ setup production` or `./server staging` - running without an environment will default to development

### To do

Add custom messages to json responses

Build a client app

### Resources

[https://thenewstack.io/make-a-restful-json-api-go/](https://thenewstack.io/make-a-restful-json-api-go/) - The overall design and a good deal of the code comes straight from this article. The author (Cory Lanou) leaves it to the reader to hook the API into a database, which is what this project sets out to do.

[http://jinzhu.me/gorm/](http://jinzhu.me/gorm/)

