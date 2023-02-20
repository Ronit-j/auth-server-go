package model

import (
	"github.com/jinzhu/gorm"
)

// User struct declaration
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username    string `json:"username"`
	Password     string `json:"password"`
}
