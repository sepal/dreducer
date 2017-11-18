package api

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"github.com/sepal/dreducer/Scanner"
)

var entityType, bundleType *graphql.Object
var queryType *graphql.Object
var Schema graphql.Schema


func setupSchema(db *Scanner.DrupalDB) {

	bundleType = graphql.NewObject(graphql.ObjectConfig{
		Name: "EntityType",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("EntityType", nil),
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	entityType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Entity",
		Description: "A drupal entity.",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Entity", nil),
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"types": &graphql.Field{
				Type: &graphql.List{
					OfType:bundleType,
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
