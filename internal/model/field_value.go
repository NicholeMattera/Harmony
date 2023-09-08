package model

import "gorm.io/gorm"

type FieldValue struct {
	gorm.Model
	EntityID	uint	`gorm:"not null"`
	Value		string	`gorm:"not null"`
	Entity		Entity
}
