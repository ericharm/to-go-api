package models

import (
    "time"
)

type Todo struct {
    ID          uint        `gorm:"primary_key" json:"id"`
    CreatedAt   time.Time   `json:"created_at"`
    UpdatedAt   time.Time   `json:"updated_at"`
    DeletedAt   *time.Time  `json:"-"`
    UserId      uint        `json:"user_id"`
    Title       string      `json:"title"`
    Description string      `json:"description"`
    Due         time.Time   `json:"due"`
    Completed   bool        `json:"completed"`
}

type Todos []Todo

