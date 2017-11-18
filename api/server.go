package api

import (
	"github.com/graphql-go/handler"
	"net/http"
	"github.com/sepal/dreducer/Scanner"
)

func InitServer(db *Scanner.DrupalDB) {
	setupSchema(db)

	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8888", nil)
}