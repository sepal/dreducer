package models

// Entity fields
type Field struct {
	table    string
	name     string
	entities []*EntityType
}

func CreateField(field_table string) Field {
	entities := make([]*EntityType, 0)

	// The name of a field is the fields table without the field_data_ prefix.
	name := field_table[11:]

	return Field{table: field_table, name: name, entities: entities}
}

func (field *Field) HasEntityType(t *EntityType) bool {
	for _, e := range field.entities {
		if e.equals(t) {
			return true
		}
	}
	return false
}

func (field *Field) AddEntityType(t *EntityType) {
	if !field.HasEntityType(t) {
		field.entities = append(field.entities, t)
	}
}

func (field *Field) Show() {
	println("Field: " + field.name)
	println("---")
	println("Belongs to:")
	for _, entity := range field.entities {
		println("- " + entity.entity.table + ":" + entity.name)

	}
	println("")
}

func (field *Field) Equals(compare *Field) bool {
	return field.table == compare.table
}