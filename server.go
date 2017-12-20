package main

import (
    "os"
    "log"
    "net/http"
    "to-go/storage"
)

func main() {
    env := "development"
    if len(os.Args) > 1 {
        env = os.Args[1]
    }

    DB := storage.InitDB(env)

    router := NewRouter()
    log.Println("Listening on port 8888")
    log.Fatal(http.ListenAndServe(":8888", router))

    DB.Close()
}


