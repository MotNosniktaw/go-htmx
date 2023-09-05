package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Message struct {
	Id      int
	Message string
}

func main() {
	router := chi.NewRouter()

	pageTmpl := template.Must(template.ParseFiles("components/page.html"))
	cardsTmpl := template.Must(template.ParseFiles("components/cards.html"))
	cardTmpl := template.Must(template.ParseFiles("components/card.html"))

	messages := []Message{
		{
			Id:      1,
			Message: "THING",
		},
		{
			Id:      2,
			Message: "THuNG",
		},
		{
			Id:      3,
			Message: "THENG",
		},
		{
			Id:      4,
			Message: "THANG",
		},
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageTmpl.Execute(w, nil)
	})

	router.Get("/cards", func(w http.ResponseWriter, r *http.Request) {
		cardsTmpl.Execute(w, messages)
	})

	router.Get("/card/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		for i := range messages {
			if messages[i].Id == id {
				cardTmpl.Execute(w, messages[i].Message)
				break
			}
		}
	})

	http.ListenAndServe(":1234", router)
}
