package models

import "time"

type User struct {
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
}
