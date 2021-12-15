package handler

import (
	"html/template"
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// w.Write([]byte("Home Page")) // plain text not html

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))

	data := map[string]string{
		"title":   "I'm Learning Golang Web",
		"content": "well, golang is super fast",
	}

	if err != nil {
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)

}
