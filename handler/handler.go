package handler

import (
	"golangweb/entity"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Home Page")) // plain text not html

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	// data := map[string]string{
	// 	"title":   "I'm Learning Golang Web",
	// 	"content": "well, golang is super fast",
	// }
	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 5}
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 11},
		{ID: 2, Name: "Ignis", Price: 190000000, Stock: 7},
		{ID: 3, Name: "Fortuner", Price: 350000000, Stock: 3},
	}

	if err != nil {
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)

}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Product Page"))
}

func ProductIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		idNumb = 0
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))

	data := map[string]int{
		"content": idNumb,
	}

	if err != nil {
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)

}
