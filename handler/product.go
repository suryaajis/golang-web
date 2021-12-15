package handler

import (
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Product Page"))
}

func ProductIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		idNumb = 0
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"))

	data := map[string]int{
		"content": idNumb,
	}

	if err != nil {
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)

}
