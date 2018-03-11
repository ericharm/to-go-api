package controllers

import (
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
    "io"
    "io/ioutil"
    "math/rand"
    "net/http"
    "regexp"
    "time"
    "github.com/gorilla/mux"
    "github.com/goware/emailx"
    "to-go/models"
    "to-go/storage"
)

type LoginAttempt struct {
    Email      string        `json:"email"`
    Password   string        `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
    login := LoginAttempt{}
    body := getBody(r)

    if err := json.Unmarshal(body, &login); err != nil {
        unprocessableEntity(w, err)
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
    signup := LoginAttempt{}
    body := getBody(r)

    if err := json.Unmarshal(body, &signup); err != nil {
        unprocessableEntity(w, err)
    }

    // validate email address format
    err := emailx.Validate(signup.Email)
    if err != nil {
        RespondWithMessage(w, "Invalid format for email.")
        return
    }

    // validate email unique
    db := storage.GetActiveDB()
    user := models.User{}
    db.Where("email = ?", signup.Email).Find(&user)

    if user.ID != 0 {
        RespondWithMessage(w, "A user with this email address already exists")
        return
    }

    // validate password
    passwordMatcher := `^[a-zA-Z0-9_.+!@#$%^&\-*]{8,32}$`
    matched, err := regexp.MatchString(passwordMatcher, signup.Password)
    if !matched {
        RespondWithMessage(w, "Password must be between 8 and 32 characters and contain alphanumeric characters or any of the following symbols: _.+!@#$%^&-*")
        return
    }

    // create user
    user.Email = signup.Email
    user.PasswordHash = hashPassword(signup.Password)
    user.AuthToken = randomString(64)
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

    body := getBody(r)

    if err := json.Unmarshal(body, &todo); err != nil {
      unprocessableEntity(w, err)
    }

    db.Save(&todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusAccepted)
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
    w.WriteHeader(http.StatusAccepted)
    if err := json.NewEncoder(w).Encode(&todo); err != nil {
        panic(err)
    }

}

// private

func getBody(r *http.Request) ([]byte) {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    return body
}

func unprocessableEntity(w http.ResponseWriter, err error) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422) // unprocessable entity
    if err := json.NewEncoder(w).Encode(err); err != nil {
        panic(err)
    }
}

var randomizer *rand.Rand // Rand for this package.

func hashPassword(password string) string {
    crypt := sha256.New()
    crypt.Write([]byte(password))
    return base64.URLEncoding.EncodeToString(crypt.Sum(nil))
}

func randomString(strlen int) string {
    const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
    result := make([]byte, strlen)
    for i := range result {
        result[i] = chars[randomizer.Intn(len(chars))]
    }
    return string(result)
}

func init() {
    randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
}

