package main

import (
    "fmt"
    "os"
    "to-go/models"
    "to-go/storage"
)

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    // open the database
    db, err := storage.ConnectGorm(env)
    if err != nil {
        panic(err)
    }

    // put all migrations here
    db.AutoMigrate(&models.Todo{})
    fmt.Println("AutoMigrating Todos.")

    db.AutoMigrate(&models.User{})
    fmt.Println("AutoMigrating Users.")

    fmt.Println("Schema up to date.")

    defer db.Close()
}

