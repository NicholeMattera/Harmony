package model

import (
	"time"
)

type Session struct {
	ID        	uint		`gorm:"primarykey"`
	IpAddress	string		`gorm:"not null;size:15"`
	Token		string		`gorm:"not null;size:128;index:idx_session_lookup"`
	LastActive	time.Time	`gorm:"not null"`
	UserId		uint		`gorm:"not null"`
	Xsrf		string		`gorm:"not null;size:36;index:idx_session_lookup"`
	User		User
}
