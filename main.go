package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/mitchellh/colorstring"
	"os"
	"regexp"
	"container/list"
)

func printError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

func main() {
	fields := list.New()

	r, _ := regexp.Compile("(field_data_.+)")

	db, err := sql.Open("mysql", "root:sws232@/drupal")

	if err != nil {
		printError(err)
	}

	processField(db, "field_data_field_job_geocomplete")

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
			fields.PushBack(table)
		}
	}

	db.Close()
}

func processField(db *sql.DB, table string) {
	rows, err := db.Query("SELECT entity_type, bundle FROM " + table)

	if err != nil {
		printError(err)
	}

	rows.Next()
	var (
		entity string
		bundle string
	)

	rows.Scan(&entity, &bundle)

	println(entity)
	println(bundle)

}
