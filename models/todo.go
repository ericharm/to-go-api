package models

import (
    "time"
    "github.com/jinzhu/gorm"
)

type Todo struct {
    gorm.Model
    Title       string      `gorm:"size:255"`
    Body        string
    Due         time.Time
    Completed   bool
}
