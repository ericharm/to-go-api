package main

import (
    "os"
    "./db"
    "./models"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    // this mapping and opening should be moved to /db
    // and reimplemented in migrate and setup scripts
    dbConfig, err := db.MapConfig()
    if err != nil {
        panic(err)
    }

    // open the database
    driver := dbConfig[env]["driver"]
    db, err := gorm.Open(driver, db.GetConnectionString(env, true))
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

