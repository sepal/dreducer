package models

// entity Fields
type Field struct {
	table     string
	Name      string        `json:"name"`
	entities  []*EntityType `json:"entities"`
	Fields    []*Field	    `json:"fields"`
}

func CreateField(field_table string) Field {
	entities := make([]*EntityType, 0)
	fields := make([]*Field, 0)

	// The Name of a field is the Fields Name without the field_data_ prefix.
	name := field_table[11:]

	return Field{table: field_table, Name: name, entities: entities, Fields:fields}
}

func (f *Field) GetName() string  {
	return f.Name
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
	println("Field: " + field.Name)
	println("---")
	println("Belongs to:")
	for _, entity := range field.entities {
		println("- " + entity.entity.Name + ":" + entity.Name)

	}
	println("")
}

func (field *Field) Equals(compare *Field) bool {
	return field.table == compare.table
}

func (f *Field) SetFieldCollection(t *EntityType)  {
	f.Fields = t.Fields
}
