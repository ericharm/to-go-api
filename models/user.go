package models

import (
    "time"
)

type User struct {
    ID             uint          `gorm:"primary_key" json:"id"`
    CreatedAt      time.Time     `json:"created_at"`
    UpdatedAt      time.Time     `json:"updated_at"`
    DeletedAt      *time.Time    `json:"-"`
    Email          string        `json:"email"`
    PasswordHash   string        `json:"-"`
    AuthToken      string        `json:"auth_token"`
}

type Users []User

