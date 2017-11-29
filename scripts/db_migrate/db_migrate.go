package main

import (
    "fmt"
    "os"
    "to-go/models"
    "to-go/db"
)

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    // open the database
    db, err := db.ConnectGorm(env)
    if err != nil {
        panic(err)
    }

    // put all migrations here
    db.AutoMigrate(&models.Todo{})
    fmt.Println("AutoMigrating Todos.")

    fmt.Println("Schema up to date.")

    defer db.Close()
}

