package main

import (
    "fmt"
    "os"
    "time"
    "to-go/models"
    "to-go/storage"
    "crypto/sha256"
    "encoding/base64"
)

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    // open the database
    DB, err := storage.ConnectGorm(env)
    if err != nil {
        panic(err)
    }
    defer DB.Close()

    wilma := models.User{
        Email: "wilma@example.com",
        PasswordHash: createPasswordHash("heckyeah"),
        AuthToken: "",
    }

    fred := models.User{
        Email: "fred@example.com",
        PasswordHash: createPasswordHash("ohhhnooo"),
        AuthToken: "",
    }

    todo1 := models.Todo{
        Title:       "Practice Go",
        Description: "Write a RESTful API",
        UserId:      1,
        Due:         time.Now(),
        Completed:   false,
    }

    todo2 := models.Todo{
        Title:       "Practice Angular",
        Description: "Write a client app to consume the Go API",
        UserId:      1,
        Due:         time.Now(),
        Completed:   false,
    }


    DB.Create(&wilma)
    DB.Create(&fred)

    fmt.Println("Added seed users.")

    DB.Create(&todo1)
    DB.Create(&todo2)

    fmt.Println("Added seed todos.")

}

func createPasswordHash(password string) string {
    crypt := sha256.New()
    crypt.Write([]byte(password))
    return base64.URLEncoding.EncodeToString(crypt.Sum(nil))
}

