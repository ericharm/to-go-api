# to-go api

### Up and running

`cd $GOPATH/src`

`git clone https://github.com/ericharm/to-go.git`

`cd to-go-api`

edit config/database.yml with your mysql credentials

`go get`

`go build`

On Mac, `export GOBIN=$GOPATH/bin` sometimes comes in handy

`./.db setup` to create a database

`./.db migrate` to create tables

`./.db seed`

`./to-go` start the server

the db&#95; scripts along with the server binary can all be run with an environment as defined in config/database.yml, e.g. `./.db setup production` or `./server staging` - running without an environment will default to development



### Resources

[https://thenewstack.io/make-a-restful-json-api-go/](https://thenewstack.io/make-a-restful-json-api-go/) - The overall design and a good deal of the code comes straight from this article. The author (Cory Lanou) leaves it to the reader to hook the API into a database, which is what this project sets out to do. I have also implemented very basic authentication.

[http://jinzhu.me/gorm/](http://jinzhu.me/gorm/)


