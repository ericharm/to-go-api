package main

import (
    "os"
    //"time"
    "log"
    "fmt"
    "html"
    "net/http"

    "./db"
    //"./models"

    "github.com/gorilla/mux"
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
    defer db.Close()

    //todo := models.Todo{
        //Title:      "My first todo",
        //Body:       "Gotta work",
        //Due:        time.Now(),
        //Completed:  false,
    //}

    //db.Create(&todo)
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    fmt.Println("Listening on port 8888")
    log.Fatal(http.ListenAndServe(":8888", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

