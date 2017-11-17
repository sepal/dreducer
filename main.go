package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/mitchellh/colorstring"
	"os"
	"regexp"
	"strings"
)

type Field struct {
	table       string
	name        string
	entities 	[]string
	bundles     []string
}

type Entity struct {
	table   string
	bundles []string
	fields  []string
}

var fields map[string]Field
var entities map[string]Entity

func printError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

func main() {
	entities = make(map[string]Entity)
	fields = make(map[string]Field)

	r, _ := regexp.Compile("(field_data_.+)")

	db, err := sql.Open("mysql", "root:sws232@/drupal")

	if err != nil {
		printError(err)
	}

	rows, err := db.Query("SHOW TABLES")

	if err != nil {
		printError(err)
	}

	for rows.Next() {
		var table string;
		err := rows.Scan(&table)
		if err != nil {
			printError(err)
		}

		if r.MatchString(table) {
			processField(db, table)
		}
	}

	//node, _ := entities["node"]
	//node.show()

	geocomplete, _ := fields["field_data_field_resume_geocomplete"]
	geocomplete.show()

	db.Close()
}

func processField(db *sql.DB, table string) {
	rows, err := db.Query("SELECT entity_type, bundle FROM " + table)

	if err != nil {
		printError(err)
	}


	field, exists := fields[table]

	if !exists {
		field = createField(table)
	}

	for rows.Next() {
		var (
			entity_type string
			bundle_name string
		)

		rows.Scan(&entity_type, &bundle_name)

		field.addBundle(bundle_name)

		entity, exists := entities[entity_type]

		if !exists {
			entity = createEntity(entity_type)
		}

		entity.addBundle(bundle_name)
		entity.addField(field)
		field.addEntity(entity)

		entities[entity_type] = entity
	}
	fields[table] = field
}

func createField(field_table string) Field {
	entities := make([]string, 0)
	bundles := make([]string, 0)

	name := field_table[11:]

	return Field{table: field_table, name:name, entities:entities, bundles: bundles}
}

func (field *Field) hasEntity(entity string) bool {
	for _, v := range field.entities {
		if v == entity {
			return true
		}
	}
	return false
}

func (field *Field) hasBundle(bundle string) bool {
	for _, v := range field.bundles {
		if v == bundle {
			return true
		}
	}
	return false
}

func (field *Field) addBundle(bundle string) {
	if !field.hasBundle(bundle) {
		field.bundles = append(field.bundles, bundle)
	}
}

func (field *Field) addEntity(entity Entity) {
	if !field.hasEntity(entity.table) {
		field.entities = append(field.entities, entity.table)
	}
}

func (field *Field) show() {
	println(field.name)
	println("---")
	println("Belongs to:")
	for _, entity_name := range field.entities {
		println("- " + entity_name)

		entity, _ := entities[entity_name]

		padding := strings.Repeat(" ", len(entity_name))
		for _, bundle := range field.bundles {
			if entity.hasBundle(bundle) {
				println("  " + padding + ":" + bundle)
			}
		}

	}


}

func createEntity(entity_name string) Entity {
	bundles := make([]string, 0)
	fields := make([]string, 0)
	return Entity{table: entity_name, bundles: bundles, fields: fields}
}

func (entity *Entity) hasBundle(bundle string) bool {
	for _, v := range entity.bundles {
		if v == bundle {
			return true
		}
	}
	return false
}

func (entity *Entity) addBundle(bundle string) {
	if !entity.hasBundle(bundle) {
		entity.bundles = append(entity.bundles, bundle)
	}
}

func (entity *Entity) hasField(field string) bool {
	for _, v := range entity.fields {
		if v == field {
			return true
		}
	}
	return false
}

func (entity *Entity) addField(field Field) {
	if !entity.hasField(field.table) {
		entity.fields = append(entity.fields, field.table)
	}
}

func (entity *Entity) show() {
	println(entity.table)
	println("---")
	println("types:")
	for _, bundle := range entity.bundles {
		println("- " + bundle)
	}

	println("fields:")
	for _, field := range entity.fields {
		f, _ := fields[field]
		println("- " + f.name)
	}
}