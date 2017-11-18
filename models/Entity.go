package models

type Entity struct {
	Name   string		 `json:"name"`
	Types  []*EntityType `json:"types"`
	Fields []*Field		 `json:"fields"`
}


func CreateEntity(entity_name string) Entity {
	bundles := make([]*EntityType, 0)
	fields := make([]*Field, 0)
	return Entity{Name: entity_name, Types: bundles, Fields: fields}
}

func (entity *Entity) HasType(name string) bool {
	for _, v := range entity.Types {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (entity *Entity) AddType(name string) {
	if !entity.HasType(name) {
		t := CreateEntityType(entity, name)
		entity.Types = append(entity.Types, &t)
	}
}

func (e *Entity) GetType(name string) (*EntityType, bool) {
	for _, t := range e.Types {
		if t.Name == name {
			return t, true
		}
	}
	return nil, false
}

func (entity *Entity) HasField(field *Field) bool {
	for _, v := range entity.Fields {
		if v.Equals(field) {
			return true
		}
	}
	return false
}

func (entity *Entity) AddField(field *Field, name string) {
	if !entity.HasField(field) {
		entity.Fields = append(entity.Fields, field)
	}
	t, _ := entity.GetType(name)
	t.addField(field)
}

func (entity *Entity) Show() {
	println("Entity: " + entity.Name)
	println("---")
	println("Types:")
	for _, t := range entity.Types {
		println("- " + t.Name)
		if len(t.fields) > 0 {
			println("  Fields:")
			for _, f := range t.fields {
				println("  - " + f.Name)
				if f.EntityRef != nil {
					for _, ef := range f.EntityRef.fields {
						println("    - " + ef.Name)
					}
				}
			}
		}
	}
	println("")
}

func (entity *Entity) Equals(compare *Entity) bool {
	return entity.Name == compare.Name
}
