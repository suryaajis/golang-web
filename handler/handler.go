package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
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

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method //untuk mengambil GET / POST / ETC

	switch method {
	case "GET":
		w.Write([]byte("ini adalah method GET"))
	case "POST":
		w.Write([]byte("ini adalah method POST"))
	default:
		http.Error(w, "Error is happening", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)

		return
	}

	http.Error(w, "Error Bad Request", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		price := r.Form.Get("price")
		stock := r.Form.Get("stock")

		data := map[string]interface{}{
			"name":  name,
			"price": price,
			"stock": stock,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data)

		return
	}

	http.Error(w, "Error Bad Request", http.StatusBadRequest)
}
