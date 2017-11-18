package models

// Entity Fields
type Field struct {
	table     string
	Name      string
	Entities  []*EntityType
	EntityRef *EntityType
}

func CreateField(field_table string) Field {
	entities := make([]*EntityType, 0)

	// The Name of a field is the Fields Name without the field_data_ prefix.
	name := field_table[11:]

	return Field{table: field_table, Name: name, Entities: entities}
}

func (field *Field) HasEntityType(t *EntityType) bool {
	for _, e := range field.Entities {
		if e.equals(t) {
			return true
		}
	}
	return false
}

func (field *Field) AddEntityType(t *EntityType) {
	if !field.HasEntityType(t) {
		field.Entities = append(field.Entities, t)
	}
}

func (field *Field) Show() {
	println("Field: " + field.Name)
	println("---")
	println("Belongs to:")
	for _, entity := range field.Entities {
		println("- " + entity.entity.Name + ":" + entity.Name)

	}
	println("")
}

func (field *Field) Equals(compare *Field) bool {
	return field.table == compare.table
}
