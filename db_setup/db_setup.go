package main

import (
    "fmt"
    "os"
    "../db"
)

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    // open the database
    sqlDb, err := db.ConnectSql(env)
    if err != nil {
        panic(err)
    }

    //validate the database exists
    err = sqlDb.Ping()
    if err != nil {
        fmt.Println(err)
        db.CreateDatabase(env)
    } else {
        fmt.Println("DB validated")
    }

    defer sqlDb.Close()
}

