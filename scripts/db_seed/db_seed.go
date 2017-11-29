package main

import (
    "fmt"
    "os"
    "time"
    "to-go/models"
    "to-go/db"
)

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    // open the database
    DB, err := db.ConnectGorm(env)
    if err != nil {
        panic(err)
    }
    defer DB.Close()

    todo1 := models.Todo{
        Title:       "Practice Go",
        Description: "Write a RESTful API",
        Due:         time.Now(),
        Completed:   false,
    }

    todo2 := models.Todo{
        Title:       "Practice Angular",
        Description: "Write a client app to consume the Go API",
        Due:         time.Now(),
        Completed:   false,
    }

    DB.Create(&todo1)
    DB.Create(&todo2)

    fmt.Println("Added seed todos.")

}

