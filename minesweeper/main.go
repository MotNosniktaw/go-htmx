package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Card struct {
	Bomb bool
}

type Row struct {
	Row []Card
}

type BoardData struct {
	Rows []Row
}

func main() {
	router := chi.NewRouter()

	pageTmpl := template.Must(template.ParseFiles("components/page.html"))
	tableTmpl := template.Must(template.ParseFiles("components/board.html"))
	cardTmpl := template.Must(template.ParseFiles("components/card.html"))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageTmpl.Execute(w, nil)
	})

	router.Get("/board", func(w http.ResponseWriter, r *http.Request) {
		cards := []Card{
			{Bomb: true},
			{Bomb: false},
			{Bomb: false},
			{Bomb: false},
			{Bomb: false},
			{Bomb: false},
			{Bomb: false},
			{Bomb: false},
			{Bomb: false},
		}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

		tableTmpl.Execute(w, cards)
	})

	router.Get("/card/true", func(w http.ResponseWriter, r *http.Request) {
		cardTmpl.Execute(w, "https://cdn-icons-png.flaticon.com/512/59/59559.png")
	})

	router.Get("/card/false", func(w http.ResponseWriter, r *http.Request) {
		cardTmpl.Execute(w, "https://cdn-icons-png.flaticon.com/512/3119/3119654.png")
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":1234", router)
}
