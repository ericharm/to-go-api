package controllers

import (
    "encoding/json"
    "net/http"
    "io"
    "io/ioutil"
    "to-go/models"
    "to-go/storage"
    "github.com/gorilla/mux"
)

func TodosIndex(w http.ResponseWriter, r *http.Request) {
    db := storage.GetActiveDB()
    user := Authenticate(r, db)

    // dry this out
    if user.ID == 0 {
        RespondWithMessage(w, "Unable to authenticate user.")
        return
    }

    todos := models.Todos{}
    db.Where("user_id = ?", user.ID).Find(&todos)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    db := storage.GetActiveDB()
    user := Authenticate(r, db)

    if user.ID == 0 {
        RespondWithMessage(w, "Unable to authenticate user.")
        return
    }

    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db.Where("user_id = ?", user.ID).Find(&todo, todoId)

    if todo.ID == 0 {
        RespondWithMessage(w, "No item was found with id " + todoId)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todo); err != nil {
        panic(err)
    }
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
    db := storage.GetActiveDB()
    user := Authenticate(r, db)

    if user.ID == 0 {
        RespondWithMessage(w, "Unable to authenticate user.")
        return
    }

    todo :=  models.Todo{ UserId: user.ID }
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

    db.Create(&todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(&todo); err != nil {
        panic(err)
    }
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
    db := storage.GetActiveDB()
    user := Authenticate(r, db)

    if user.ID == 0 {
        RespondWithMessage(w, "Unable to authenticate user.")
        return
    }

    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db.Where("user_id = ?", user.ID).Find(&todo, todoId)

    if todo.ID == 0 {
        RespondWithMessage(w, "No items with id " + todoId + " found for this user.")
        return
    }

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

    db.Save(&todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(&todo); err != nil {
        panic(err)
    }
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
    db := storage.GetActiveDB()
    user := Authenticate(r, db)

    if user.ID == 0 {
        RespondWithMessage(w, "Unable to authenticate user.")
        return
    }

    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db.Where("user_id = ?", user.ID).Find(&todo, todoId)

    if todo.ID == 0 {
        RespondWithMessage(w, "No items with id " + todoId + " found for this user.")
        return
    }

    db.Delete(&todo, todoId)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(&todo); err != nil {
        panic(err)
    }

}

