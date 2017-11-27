package main

import (
    "fmt"
    "os"
    "../models"
    "../db"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    dbConfig, err := db.MapConfig()
    if err != nil {
        panic(err)
    }

    // open the database
    envConfig := dbConfig[env]
    db, err := gorm.Open(envConfig["driver"], db.GetConnectionString(env, true))
    if err != nil {
        panic(err)
    }

    // put all migrations here
    db.AutoMigrate(&models.Todo{})
    fmt.Println("AutoMigrating Todos.")

    fmt.Println("Schema up to date.")

    defer db.Close()
}


