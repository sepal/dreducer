package models

type Entity struct {
	table  string
	types  []*EntityType
	fields []*Field
}


func CreateEntity(entity_name string) Entity {
	bundles := make([]*EntityType, 0)
	fields := make([]*Field, 0)
	return Entity{table: entity_name, types: bundles, fields: fields}
}

func (entity *Entity) HasType(name string) bool {
	for _, v := range entity.types {
		if v.name == name {
			return true
		}
	}
	return false
}

func (entity *Entity) AddType(name string) {
	if !entity.HasType(name) {
		t := CreateEntityType(entity, name)
		entity.types = append(entity.types, &t)
	}
}

func (e *Entity) GetType(name string) (*EntityType, bool) {
	for _, t := range e.types {
		if t.name == name {
			return t, true
		}
	}
	return nil, false
}

func (entity *Entity) HasField(field *Field) bool {
	for _, v := range entity.fields {
		if v.Equals(field) {
			return true
		}
	}
	return false
}

func (entity *Entity) AddField(field *Field, name string) {
	if !entity.HasField(field) {
		entity.fields = append(entity.fields, field)
		t, _ := entity.GetType(name)
		t.addField(field)
	}
}

func (entity *Entity) Show() {
	println("Entity: " + entity.table)
	println("---")
	println("types:")
	for _, t := range entity.types {
		println("- " + t.name)
		if len(t.fields) > 0 {
			println("  fields:")
			for _, f := range t.fields {
				println("  - " + f.name)
			}
		}
	}
	println("")
}

func (entity *Entity) Equals(compare *Entity) bool {
	return entity.table == compare.table
}
