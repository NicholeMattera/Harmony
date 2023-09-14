package store

type Composition struct {
	StandardAudit
	DefaultPermissionFlags	PermissionFlag	`gorm:"not null,default:0"`
	Name 					string			`gorm:"not null"`
}

func (s *storeLayer) CreateComposition(createdBy, name string, defaultPermissionFlag PermissionFlag) (Composition, error) {
    return Composition{}, nil
}

func (s *storeLayer) GetAllCompositions() ([]Composition, error) {
	return []Composition{}, nil
}
