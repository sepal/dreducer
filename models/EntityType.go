package models

type EntityType struct {
	entity *Entity
	Name   string
	fields []*Field
}

func CreateEntityType(e *Entity, name string) (EntityType) {
	fields := make([]*Field, 0)
	return EntityType{entity:e, Name: name, fields: fields}
}

func (t *EntityType) equals(c *EntityType) bool {
	return t.Name == c.Name && t.entity.Equals(c.entity)
}

func (t *EntityType) hasFields(field *Field)  bool{
	for _, f := range t.fields {
		if f.Equals(field) {
			return true
		}
	}

	return false
}

func (t *EntityType) addField(f *Field) {
	if !t.hasFields(f) {
		t.fields = append(t.fields, f)
	}
}
func (t *EntityType) Show() {
	println("Entity: " + t.entity.Name + ":" + t.Name)
	println("---")
	println("Fields: ")
	for _, f := range t.fields {
		println("- " + f.Name)
		if f.EntityRef != nil {
			for _, ef := range f.EntityRef.fields {
				println("  - " + ef.Name)
			}
		}
	}
}