package main

import (
    "fmt"
    "os"
    "./db"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func create(envConfig map[string]string) {
    driver := envConfig["driver"]
    sqlDB, err := sql.Open(driver, db.GetConnectionString(os.Args[1], false))
    if err != nil {
        fmt.Println(err)
    }
    defer sqlDB.Close()

    _, err = sqlDB.Exec("CREATE DATABASE " + envConfig["database"])
    if err == nil {
        fmt.Printf("DB '%s' Created.\n", envConfig["database"])
    } else {
        panic(err)
    }

    _, err = sqlDB.Exec("USE " + envConfig["database"])
    if err == nil {
        fmt.Printf("Using %s\n", envConfig["database"])
    } else {
        panic(err)
    }
}

func main() {
    env := os.Args[1]

    dbConfig, err := db.MapConfig()
    if err != nil {
        panic(err)
    }

    // open the database
    driver := dbConfig[env]["driver"]
    sqlDB, err := sql.Open(driver, db.GetConnectionString(env, true))
    if err != nil {
        panic(err)
    }

    //validate the database exists
    err = sqlDB.Ping()
    if err != nil {
        fmt.Println(err)
        // create the db if it doesn't exist
        create(dbConfig[env])
    } else {
        fmt.Println("DB validated")
    }

    defer sqlDB.Close()
}

