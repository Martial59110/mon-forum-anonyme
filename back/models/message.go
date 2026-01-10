package models

import "time"

type Message struct {
	ID        int       `json:"id" db:"id"`
	Pseudonym string    `json:"pseudonym" db:"pseudonym"`
	Content   string    `json:"content" db:"content"`
	Avatar    string    `json:"avatar" db:"avatar"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
