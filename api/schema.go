package api

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/sepal/dreducer/Scanner"
	"github.com/sepal/dreducer/models"
	"golang.org/x/net/context"
	"errors"
)

var entityType, bundleType, fieldType, fieldFieldType *graphql.Object
var queryType *graphql.Object
var Schema graphql.Schema

func tableId(obj interface{}, info graphql.ResolveInfo, ctx context.Context)(string, error) {
	e:= obj.(models.Table)
	return e.GetName(), nil
}

func setupSchema(db *Scanner.DrupalDB) {
	fieldFieldType  = graphql.NewObject(graphql.ObjectConfig{
		Name: "FieldField",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("FieldField", tableId),
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	fieldType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Field",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Field", tableId),
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"fields": &graphql.Field{
				Type: &graphql.List{
					OfType: fieldFieldType,
				},
			},
		},
	})

	bundleType = graphql.NewObject(graphql.ObjectConfig{
		Name: "EntityType",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("EntityType", tableId),
			"type": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					t := p.Source.(*models.EntityType)
					return t.Name, nil
				},
			},
			"fields": &graphql.Field{
				Type: &graphql.List{
					OfType: fieldType,
				},
			},
		},
	})

	entityType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "entity",
		Description: "A drupal entity.",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("entity", tableId),
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"types": &graphql.Field{
				Type: &graphql.List{
					OfType: bundleType,
				},
				Args: graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					e := p.Source.(*models.Entity)
					bundle, exists := p.Args["type"].(string)

					if exists {
						t, _ := e.GetType(bundle)
						types := make([]*models.EntityType, 1)
						types[0] = t
						return types, nil
					}

					return e.Types, nil
				},
			},
			"fields": &graphql.Field{
				Type: &graphql.List{
					OfType: fieldType,
				},
			},
		},
	})

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"entities": &graphql.Field{
				Type: graphql.NewList(entityType),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Description: "If ommitted, returns all entities " +
							"otherwise the entity with the provided name.",
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, exists := p.Args["name"].(string)

					if exists {
						e, _ := db.GetEntity(name)
						entities := make([]*models.Entity, 1)
						entities[0] = e
						return entities, nil
					}

					e := db.All()
					return e, nil
				},
			},
			"entityType": &graphql.Field{
				Type: bundleType,
				Args: graphql.FieldConfigArgument{
					"entity": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					e, exists := p.Args["entity"].(string)

					if !exists {
						return nil, errors.New("missing entity name argument")
					}

					b, exists := p.Args["type"].(string)

					if !exists {
						return nil, errors.New("missing entity type name argument")
					}

					t, exists := db.GetEntityType(e, b)

					if !exists {
						return nil, errors.New("entity type not found")
					}

					return t, nil
				},
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

}
