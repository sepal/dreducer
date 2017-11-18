package models

type EntityType struct {
	entity *Entity
	name string
	fields []*Field
}

func CreateEntityType(e *Entity, name string) (EntityType) {
	fields := make([]*Field, 0)
	return EntityType{entity:e, name: name, fields: fields}
}

func (t *EntityType) equals(c *EntityType) bool {
	return t.name == c.name && t.entity.Equals(c.entity)
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