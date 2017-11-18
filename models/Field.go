package models

import "strings"

// Entity fields
type Field struct {
	table    string
	name     string
	entities []*Entity
	bundles  []string
}

func CreateField(field_table string) Field {
	entities := make([]*Entity, 0)
	bundles := make([]string, 0)

	// The name of a field is the fields table without the field_data_ prefix.
	name := field_table[11:]

	return Field{table: field_table, name: name, entities: entities, bundles: bundles}
}

func (field *Field) HasEntity(entity *Entity) bool {
	for _, v := range field.entities {
		if v.Equals(entity) {
			return true
		}
	}
	return false
}

func (field *Field) HasBundle(bundle string) bool {
	for _, v := range field.bundles {
		if v == bundle {
			return true
		}
	}
	return false
}

func (field *Field) AddBundle(bundle string) {
	if !field.HasBundle(bundle) {
		field.bundles = append(field.bundles, bundle)
	}
}

func (field *Field) AddEntity(entity *Entity) {
	if !field.HasEntity(entity) {
		field.entities = append(field.entities, entity)
	}
}

func (field *Field) Show() {
	println("Field: " + field.name)
	println("---")
	println("Belongs to:")
	for _, entity := range field.entities {
		println("- " + entity.table)

		padding := strings.Repeat(" ", len(entity.table))
		for _, bundle := range field.bundles {
			if entity.HasBundle(bundle) {
				println("   " + padding + ":" + bundle)
			}
		}

	}
	println("")
}

func (field *Field) Equals(compare *Field) bool {
	return field.table == compare.table
}