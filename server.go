package main

import (
    "os"
    "./db"
    "./models"
    "time"
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

    todo := models.Todo{
        Title:      "My first todo",
        Body:       "Gotta work",
        Due:        time.Now(),
        Completed:  false,
    }

    db.Create(&todo)

    defer db.Close()

}

