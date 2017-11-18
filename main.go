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

	//err = api.UpdateSchema(&db)
	//
	//if err != nil {
	//	printError(err)
	//}

	api.InitServer(&db)
}
