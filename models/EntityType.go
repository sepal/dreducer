package models

type EntityType struct {
	entity *Entity
	name string
}

func CreateEntityType(e *Entity, name string) (EntityType) {
	return EntityType{entity:e, name: name}
}

func (t *EntityType) equals(c *EntityType) bool {
	return t.name == c.name && t.entity.Equals(c.entity)
}
