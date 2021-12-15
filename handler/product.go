package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Product Page"))
}

func ProductIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		w.Write([]byte("Product Page"))
	} else {
		fmt.Fprintf(w, "Product page by id: %d", idNumb)
	}

}
