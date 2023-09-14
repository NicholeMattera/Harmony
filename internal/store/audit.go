package store

import (
	"gorm.io/gorm"
)

type StandardAudit struct {
	gorm.Model
	CreatedBy	string	`gorm:"not null"`
	UpdatedBy	string
	DeletedBy	string
}
