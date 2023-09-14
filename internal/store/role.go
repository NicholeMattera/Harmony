package store

import "github.com/Thor-x86/nullable"

type Role struct {
	StandardAudit
	LockSessionToIp		bool			`gorm:"not null;default:false"`
	MultipleSessions	bool			`gorm:"not null;default:true"`
	Name				string			`gorm:"not null"`
	SessionTimeout		nullable.Uint64
}

func (s *storeLayer) CreateRole(createdBy, lockSessionToIp, multipleSessions bool, name string, sessionTimeout *uint64) (Role, error) {
    return Role{}, nil
}

func (s *storeLayer) GetAllRoles() ([]Role, error) {
	return []Role{}, nil
}
