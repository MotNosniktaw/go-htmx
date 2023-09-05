package main

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	count := 0

	router := chi.NewRouter()

	pageTmpl := template.Must(template.ParseFiles("components/page.html"))
	counterTmpl := template.Must(template.ParseFiles("components/poll.html"))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageTmpl.Execute(w, count)
	})

	router.Get("/poll", func(w http.ResponseWriter, r *http.Request) {
		count++
		counterTmpl.Execute(w, count)
	})

	http.ListenAndServe(":1234", router)
}
