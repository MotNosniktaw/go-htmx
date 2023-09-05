package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type Todo struct {
	Id    int
	Title string
	Done  bool
}

type TodoTableData struct {
	Todos []Todo
}

type TodoPageData struct {
	PageTitle string
}

func main() {
	router := chi.NewRouter()

	pageTmpl := template.Must(template.ParseFiles("components/page.html"))
	tableTmpl := template.Must(template.ParseFiles("components/table.html"))
	rowTmpl := template.Must(template.ParseFiles("components/row.html"))

	tableData := TodoTableData{
		Todos: []Todo{
			{
				Id:    1,
				Title: "Thing",
				Done:  true,
			},
			{
				Id:    2,
				Title: "Theng",
				Done:  true,
			},
			{
				Id:    3,
				Title: "Thung",
				Done:  false,
			},
			{
				Id:    4,
				Title: "Thang",
				Done:  false,
			},
			{
				Id:    5,
				Title: "Thing",
				Done:  true,
			},
			{
				Id:    6,
				Title: "Theng",
				Done:  true,
			},
			{
				Id:    7,
				Title: "Thung",
				Done:  false,
			},
			{
				Id:    8,
				Title: "Thang",
				Done:  false,
			},
			{
				Id:    9,
				Title: "Thing",
				Done:  true,
			},
			{
				Id:    10,
				Title: "Theng",
				Done:  true,
			},
			{
				Id:    11,
				Title: "Thung",
				Done:  false,
			},
			{
				Id:    12,
				Title: "Thang",
				Done:  false,
			},
			{
				Id:    13,
				Title: "Thing",
				Done:  true,
			},
			{
				Id:    14,
				Title: "Theng",
				Done:  true,
			},
			{
				Id:    15,
				Title: "Thung",
				Done:  false,
			},
			{
				Id:    16,
				Title: "Thang",
				Done:  false,
			},
			{
				Id:    17,
				Title: "Thing",
				Done:  true,
			},
			{
				Id:    18,
				Title: "Theng",
				Done:  true,
			},
			{
				Id:    3,
				Title: "Thung",
				Done:  false,
			},
			{
				Id:    4,
				Title: "Thang",
				Done:  false,
			},
			{
				Id:    1,
				Title: "Thing",
				Done:  true,
			},
			{
				Id:    2,
				Title: "Theng",
				Done:  true,
			},
			{
				Id:    3,
				Title: "Thung",
				Done:  false,
			},
			{
				Id:    4,
				Title: "Thang",
				Done:  false,
			},
			{
				Id:    1,
				Title: "Thing",
				Done:  true,
			},
			{
				Id:    2,
				Title: "Theng",
				Done:  true,
			},
			{
				Id:    3,
				Title: "Thung",
				Done:  false,
			},
			{
				Id:    4,
				Title: "Thang",
				Done:  false,
			},
			{
				Id:    1,
				Title: "Thing",
				Done:  true,
			},
			{
				Id:    2,
				Title: "Theng",
				Done:  true,
			},
			{
				Id:    3,
				Title: "Thung",
				Done:  false,
			},
			{
				Id:    4,
				Title: "Thang",
				Done:  false,
			},
		},
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pageData := TodoPageData{
			PageTitle: "TODO - HTMX Lazy List",
		}

		pageTmpl.Execute(w, pageData)
	})

	router.Get("/todos", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		n := 100 + rand.Intn(400)
		time.Sleep(time.Duration(n) * time.Millisecond)

		tableTmpl.Execute(w, tableData)
	})

	router.Get("/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			fmt.Println("Could get valid id from path")
			w.WriteHeader(400)
			w.Write([]byte("invalid id"))
		}

		for i := range tableData.Todos {
			if tableData.Todos[i].Id == id {
				rand.Seed(time.Now().UnixNano())
				n := 100 + rand.Intn(400)
				time.Sleep(time.Duration(n) * time.Millisecond)

				rowTmpl.Execute(w, tableData.Todos[i])
				break
			}
		}
	})

	http.ListenAndServe(":1234", router)
}
