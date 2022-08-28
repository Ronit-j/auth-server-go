package model

import (
	"github.com/jinzhu/gorm"
)

// User struct declaration
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Salt     string
	Verifier string
}
