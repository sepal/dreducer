package models

type EntityType struct {
	entity *Entity  `json:"entity"`
	Name   string   `json:"name"`
	Fields []*Field `json:"fields"`
}

func CreateEntityType(e *Entity, name string) (EntityType) {
	fields := make([]*Field, 0)
	return EntityType{entity: e, Name: name, Fields: fields}
}

func (t *EntityType) equals(c *EntityType) bool {
	return t.Name == c.Name && t.entity.Equals(c.entity)
}

func (t *EntityType) hasFields(field *Field) bool {
	for _, f := range t.Fields {
		if f.Equals(field) {
			return true
		}
	}

	return false
}

func (t *EntityType) addField(f *Field) {
	if !t.hasFields(f) {
		t.Fields = append(t.Fields, f)
	}
}
func (t *EntityType) Show() {
	println("entity: " + t.entity.Name + ":" + t.Name)
	println("---")
	println("Fields: ")
	for _, f := range t.Fields {
		println("- " + f.Name)
		if f.entityRef != nil {
			for _, ef := range f.entityRef.Fields {
				println("  - " + ef.Name)
			}
		}
	}
}