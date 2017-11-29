package controllers

import (
    "encoding/json"
    "net/http"
    "io"
    "io/ioutil"
    "to-go/models"
    "to-go/db"
    "github.com/gorilla/mux"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    db_ := db.GetActiveDB()
    todos := models.Todos{}
    db_.Find(&todos)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    db_ := db.GetActiveDB()
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db_.Find(&todo, todoId)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todo); err != nil {
        panic(err)
    }
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
    db_ := db.GetActiveDB()
    todo :=  models.Todo{}
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    db_.Create(&todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(&todo); err != nil {
        panic(err)
    }
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
    db_ := db.GetActiveDB()
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db_.Find(&todo, todoId)

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    db_.Save(&todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(&todo); err != nil {
        panic(err)
    }
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
    db_ := db.GetActiveDB()
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db_.Find(&todo, todoId)
    db_.Delete(&todo)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(&todo); err != nil {
        panic(err)
    }

}

