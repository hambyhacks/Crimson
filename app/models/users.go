package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"username" validate:"username,min=4,max=20,required"`
	Email     string    `json:"email" validate:"email,required"`
	Password  Password  `json:"-"`
	Activated bool      `json:"activated"`
}

type Password struct {
	Plaintext string
	Hash      []byte
}
