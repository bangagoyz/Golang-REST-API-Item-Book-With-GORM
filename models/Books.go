package models

import (
	"time"
)

type Book struct {
	ID        uint      `json:"id"`
	BookName  string    `json:"name_book"`
	Author    string    `json:"author"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
