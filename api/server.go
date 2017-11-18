package api

import (
	"github.com/graphql-go/handler"
	"net/http"
	"github.com/sepal/dreducer/Scanner"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
	"encoding/json"
	"io/ioutil"
	"os"
	"errors"
)

func UpdateSchema(db *Scanner.DrupalDB) error {
	setupSchema(db)

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: testutil.IntrospectionQuery,
	})

	if result.HasErrors() {
		return errors.New("error while generating schema")
	}
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./visualizer/data/schema.json", b, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func InitServer(db *Scanner.DrupalDB) {
	setupSchema(db)

	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8888", nil)
}
