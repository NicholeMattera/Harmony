package model

import (
	"database/sql"
)

type User struct {
	StandardAudit
	FirstName		string			`gorm:"not null;size:64"`
	Email			string			`gorm:"not null;size:254"`
	LastLoggedInOn	sql.NullTime
	LastName 		string			`gorm:"not null;size:64"`
	Username		string			`gorm:"uniqueIndex;not null;size:64"`
	Password		string			`gorm:"not null;size:64"`
	RoleID			uint			`gorm:"not null"`
	Role 			Role
}
