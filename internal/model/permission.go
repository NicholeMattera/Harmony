package model

type PermissionFlag uint8

const (
	Read		PermissionFlag = 1
	Write		PermissionFlag = 2
)

type Permission struct {
	Flags 			PermissionFlag	`gorm:"not null"`
	CompositionID	uint			`gorm:"not null"`
	RoleID			uint			`gorm:"not null"`
	Composition		Composition
	Role			Role
}
