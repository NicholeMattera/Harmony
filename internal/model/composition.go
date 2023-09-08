package model

type Composition struct {
	StandardAudit
	DefaultPermissionFlags	PermissionFlag	`gorm:"not null,default:0"`
	Name 					string			`gorm:"not null"`
}
