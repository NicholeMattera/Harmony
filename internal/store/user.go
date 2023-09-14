package store

import (
	"database/sql"
)

type User struct {
	StandardAudit
	Email			string			`gorm:"not null;size:254"`
	FirstName		string			`gorm:"not null;size:64"`
	LastLoggedInOn	sql.NullTime
	LastName 		string			`gorm:"not null;size:64"`
	Username		string			`gorm:"uniqueIndex;not null;size:64"`
	Password		string			`gorm:"not null;size:64"`
	RoleID			uint			`gorm:"not null"`
	Role 			Role
}

func (s *storeLayer) CreateUser(createdBy, email, firstName, lastName, username, password string, roleId uint) (User, error) {
    return User{}, nil
}

func (s *storeLayer) GetAllUsers() ([]User, error) {
	return []User{}, nil
}
