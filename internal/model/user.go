package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `grom:"unique"`
	PasswordHash string
	Salt         string
	Session      string `grom:"uniqueIndex"`
}
