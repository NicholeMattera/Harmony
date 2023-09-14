package store

type FieldType string

const (
	Association	FieldType = "association"
	Checkbox	FieldType = "checkbox"
	Color		FieldType = "color"
	Date		FieldType = "date"
	DateTime	FieldType = "date-time"
	Dropdown	FieldType = "dropdown"
	Email		FieldType = "email"
	File		FieldType = "file"
	LongText 	FieldType = "long-text"
	Number		FieldType = "number"
	Phone		FieldType = "phone"
	ShortText   FieldType = "short-text"
	Time		FieldType = "time"
	URL			FieldType = "url"
)

type Field struct {
	StandardAudit
	CompositionID	uint					`gorm:"not null"`
	DefaultValue	string					`gorm:"not null"`
	Metadata		map[string]interface{}	`gorm:"serializer:json"`
	Name			string					`gorm:"not null"`
	Position		uint8					`gorm:"not null;default:0"`
	Required		bool					`gorm:"not null;default:false"`
	Type			FieldType				`gorm:"not null"`
	Composition		Composition
}

func (s *storeLayer) CreateField(compositionID uint, createdBy, defaultValue string, fieldType FieldType, metadata map[string]interface{}, name string, position uint8, required bool) (Field, error) {
    return Field{}, nil
}
