package store

import "github.com/Thor-x86/nullable"

type Entity struct {
	StandardAudit
	CompositionID	uint			`gorm:"not null"`
	Name			string			`gorm:"not null"`
	ParentEntityID	nullable.Uint
	Slug			string			`gorm:"uniqueIndex;not null;size:512"`
	Composition		Composition
	Parent			*Entity			`gorm:"foreignKey:ParentEntityID"`
}

func (s *storeLayer) CreateEntity(compositionID uint, createdBy, name string, parentEntityId *uint, slug string) (Entity, error) {
    return Entity{}, nil
}

func (s *storeLayer) GetAllEntities() ([]Entity, error) {
	return []Entity{}, nil
}
