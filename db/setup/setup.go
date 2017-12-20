package main

import (
    "fmt"
    "os"
    "to-go/storage"
)

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    // open the database
    db, err := storage.ConnectSql(env)
    if err != nil {
        panic(err)
    }

    //validate the database exists
    err = db.Ping()
    if err != nil {
        fmt.Println(err)
        storage.CreateDatabase(env)
    } else {
        fmt.Println("DB validated")
    }

    defer db.Close()
}

