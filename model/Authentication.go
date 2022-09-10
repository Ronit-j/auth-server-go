package model

import (
	"github.com/jinzhu/gorm"
)

// User struct declaration
type Authentication_User struct {
	gorm.Model
	Username    string `json:"username"`
	SessionKey  string `json:"SessionKey"`
}
