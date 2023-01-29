package model

import "gorm.io/gorm"

// use gorm to create a model of user, with email, password, username
type User struct {
	gorm.Model
	Email    string
	Password string
	Username string
}
