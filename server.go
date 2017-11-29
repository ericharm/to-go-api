package main

import (
    "os"
    "log"
    "fmt"
    "html"
    "net/http"
    "encoding/json"

    "./db"
    "./models"

    "github.com/jinzhu/gorm"
    "github.com/gorilla/mux"
)

var db_ *gorm.DB

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    // open the database
    gormDb, err := db.ConnectGorm(env)
    if err != nil {
        panic(err)
    }

    defer gormDb.Close()
    db_ = gormDb

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/todos", TodoIndex)
    router.HandleFunc("/todos/{todoId}", TodoShow)

    fmt.Println("Listening on port 8888")
    log.Fatal(http.ListenAndServe(":8888", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := []models.Todo{}
    db_.Find(&todos)
    json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db_.Find(&todo, todoId)
    json.NewEncoder(w).Encode(todo)
}

