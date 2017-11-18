package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/colorstring"
	"os"
	"github.com/sepal/dreducer/Scanner"
	"github.com/sepal/dreducer/api"
)

func printError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

func main() {
	db, err := Scanner.CreateDrupalDB("drupal", "", "root", "sws232")

	if err != nil {
		printError(err)
	}

	api.InitServer(&db)

	//name := "profile2"
	//
	//entity, exists := db.GetEntity(name)
	////entity, exists := db.GetEntityType(name, "resume")
	//if !exists {
	//	fmt.Printf("Entity %s does not exist.\n", name)
	//	os.Exit(1)
	//}
	//entity.Show()
}
