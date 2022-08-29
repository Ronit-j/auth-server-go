package model

import (
	"github.com/jinzhu/gorm"
)

// User struct declaration
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username    string `json:"username"`
	Salt     string `json:"salt"`
	Verifier string `json:"verifier"`
}
