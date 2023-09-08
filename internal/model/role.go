package model

type Role struct {
	StandardAudit
	LockSessionToIp		bool	`gorm:"not null;default:false"`
	MultipleSessions	bool	`gorm:"not null;default:true"`
	Name				string	`gorm:"not null"`
	SessionTimeout		uint64
}
