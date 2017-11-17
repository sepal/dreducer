package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/mitchellh/colorstring"
	"os"
	"regexp"
)

type Field struct {
	table       string
	name        string
	entity_type []Entity
	bundle      []string
}

type Entity struct {
	table   string
	bundles []string
	fields  []Field
}

var entities map[string]Entity

func printError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

func main() {
	entities = make(map[string]Entity)

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

	for k, entity := range entities {
		println(k)
		for _, bundle := range entity.bundles {
			println("- " + bundle)
		}
	}

	db.Close()
}

func processField(db *sql.DB, table string) {
	rows, err := db.Query("SELECT entity_type, bundle FROM " + table)

	if err != nil {
		printError(err)
	}

	for rows.Next() {
		var (
			entity_type string
			bundle_name string
		)

		rows.Scan(&entity_type, &bundle_name)

		if val, ok := entities[entity_type]; ok {
			val.addBundle(bundle_name)
			entities[entity_type] = val
		} else {
			bundles := make([]string, 1)
			bundles[0] = bundle_name
			entity := Entity{table: entity_type, bundles: bundles}

			entities[entity_type] = entity
		}
	}
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