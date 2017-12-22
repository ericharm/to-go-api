package controllers

import (
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
    "github.com/gorilla/mux"
    "io"
    "io/ioutil"
    "net/http"
    "to-go/models"
    "to-go/storage"
    "math/rand"
    "time"
)

var r *rand.Rand // Rand for this package.

type LoginAttempt struct {
    Email      string        `json:"email"`
    Password   string        `json:"password"`
}

// replace this with controller.MessageResponse
type LoginFailResponse struct {
    Message    string `json:"message"`
}

func Login(w http.ResponseWriter, r *http.Request) {
    login := LoginAttempt{}
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &login); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    hash := hashPassword(login.Password)
    db := storage.GetActiveDB()
    user := models.User{}
    db.Where("email = ? AND password_hash = ?", login.Email, hash).First(&user)

    if user.ID == 0 {
        RespondWithMessage(w, "Invalid email address or password.")
        return
    }

    user.AuthToken = randomString(64)
    db.Save(&user)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(&user); err != nil {
        panic(err)
    }
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
    db := storage.GetActiveDB()
    user := models.User{}
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &user); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    db.Create(&user)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(&user); err != nil {
        panic(err)
    }
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
    db := storage.GetActiveDB()
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db.Find(&todo, todoId)

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

func UserDelete(w http.ResponseWriter, r *http.Request) {
    db := storage.GetActiveDB()
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    todo := models.Todo{}
    db.Find(&todo, todoId)
    db.Delete(&todo, todoId)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(&todo); err != nil {
        panic(err)
    }

}

func hashPassword(password string) string {
    crypt := sha256.New()
    crypt.Write([]byte(password))
    return base64.URLEncoding.EncodeToString(crypt.Sum(nil))
}

func randomString(strlen int) string {
    const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
    result := make([]byte, strlen)
    for i := range result {
        result[i] = chars[r.Intn(len(chars))]
    }
    return string(result)
}

// replace this with message response from base controller module
func newLoginFailResponse() LoginFailResponse {
    response := LoginFailResponse{}
    response.Message = ""
    return response
}

func init() {
    r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

