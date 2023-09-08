package model

type Entity struct {
	StandardAudit
	CompositionID	uint		`gorm:"not null"`
	Name			string		`gorm:"not null"`
	ParentEntityID	uint
	Slug			string		`gorm:"uniqueIndex;not null;size:512"`
	Composition		Composition
	Parent			*Entity		`gorm:"foreignKey:ParentEntityID"`
}
