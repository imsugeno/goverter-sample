package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	gen "github.com/imsugeno/query-generator-sample/api/gen"
	convertergen "github.com/imsugeno/query-generator-sample/converter/gen"
	"github.com/imsugeno/query-generator-sample/handler"
	"github.com/imsugeno/query-generator-sample/query"
)

func main() {
	q := &query.InMemoryTaskQuery{}
	c := &convertergen.ConverterImpl{}
	h := handler.New(q, c)

	r := chi.NewRouter()
	gen.HandlerFromMux(h, r)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
