package api

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/sepal/dreducer/Scanner"
)

var entityType, bundleType, fieldType *graphql.Object
var queryType *graphql.Object
var Schema graphql.Schema


func setupSchema(db *Scanner.DrupalDB) {

	fieldType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Field",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Field", nil),
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	bundleType = graphql.NewObject(graphql.ObjectConfig{
		Name: "EntityType",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("EntityType", nil),
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"fields": &graphql.Field{
				Type: &graphql.List{
					OfType: fieldType,
				},
			},
		},
	})

	entityType = graphql.NewObject(graphql.ObjectConfig{
		Name: "entity",
		Description: "A drupal entity.",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("entity", nil),
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"types": &graphql.Field{
				Type: &graphql.List{
					OfType:bundleType,
				},
			},
			"fields": &graphql.Field{
				Type: &graphql.List{
					OfType:fieldType,
				},
			},
		},
	})

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"entity": &graphql.Field{
				Type: entityType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Description: "If ommitted, returns all entities " +
							"otherwise the entity with the provided name.",
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					name, _ := p.Args["name"].(string)
					e, _ := db.GetEntity(name)
					return e, nil
				},
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

}
