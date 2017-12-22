package controllers

import (
    "encoding/json"
    "net/http"
    "to-go/models"
    "strings"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type MessageResponse struct {
    Status      string `json:"status"`
    Message     string `json:"message"`
}

func RespondWithMessage(w http.ResponseWriter, message string) {
    response := MessageResponse{
        Status: "fail", Message: message,
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(404)
    if err := json.NewEncoder(w).Encode(&response); err != nil {
        panic(err)
    }
}

func Authenticate(r *http.Request, db *gorm.DB) *models.User {
    authToken := strings.Join(r.Header["X-Auth-Token"], "")
    user := models.User{}
    db.Where("auth_token = ?", authToken).First(&user)
    return &user
}


