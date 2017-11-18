package main

import "database/sql"
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/colorstring"
	"github.com/sepal/dreducer/models"
	"os"
	"regexp"
)

var fields map[string]*models.Field
var entities map[string]*models.Entity

func printError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

func main() {
	entities = make(map[string]*models.Entity)
	fields = make(map[string]*models.Field)

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
		var table string
		err := rows.Scan(&table)
		if err != nil {
			printError(err)
		}

		if r.MatchString(table) {
			processField(db, table)
		}
	}

	node, _ := entities["node"]
	node.Show()

	geocomplete, _ := fields["field_data_field_resume_geocomplete"]
	geocomplete.Show()

	db.Close()
}

func processField(db *sql.DB, table string) {
	rows, err := db.Query("SELECT entity_type, bundle FROM " + table)

	if err != nil {
		printError(err)
	}

	field, exists := fields[table]

	if !exists {
		f := models.CreateField(table)
		field = &f
	}

	for rows.Next() {
		var (
			entity_type string
			bundle_name string
		)

		rows.Scan(&entity_type, &bundle_name)

		entity, exists := entities[entity_type]

		if !exists {
			e := models.CreateEntity(entity_type)
			entity = &e
		}

		t, exists := entity.GetType(bundle_name)

		if !exists {
			entity.AddType(bundle_name)
			t, _ = entity.GetType(bundle_name)
		}

		field.AddEntityType(t)

		entity.AddField(field, bundle_name)

		entities[entity_type] = entity
	}
	fields[table] = field
}
