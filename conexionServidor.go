package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/Hola mundo", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola mundo desde go")
}
