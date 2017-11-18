package models

type Entity struct {
	table   string
	bundles []string
	fields  []*Field
}


func CreateEntity(entity_name string) Entity {
	bundles := make([]string, 0)
	fields := make([]*Field, 0)
	return Entity{table: entity_name, bundles: bundles, fields: fields}
}

func (entity *Entity) HasBundle(bundle string) bool {
	for _, v := range entity.bundles {
		if v == bundle {
			return true
		}
	}
	return false
}

func (entity *Entity) AddBundle(bundle string) {
	if !entity.HasBundle(bundle) {
		entity.bundles = append(entity.bundles, bundle)
	}
}

func (entity *Entity) HasField(field *Field) bool {
	for _, v := range entity.fields {
		if v.Equals(field) {
			return true
		}
	}
	return false
}

func (entity *Entity) AddField(field *Field) {
	if !entity.HasField(field) {
		entity.fields = append(entity.fields, field)
	}
}

func (entity *Entity) Show() {
	println("Entity: " + entity.table)
	println("---")
	println("types:")
	for _, bundle := range entity.bundles {
		println("- " + bundle)
	}

	println("fields:")
	for _, f := range entity.fields {
		println("- " + f.name)
	}
	println("")
}

func (entity *Entity) Equals(compare *Entity) bool {
	return entity.table == compare.table
}