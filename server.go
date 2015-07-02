package main

import (
	"net/http"
	"fmt"
//	"os"
)

var listenAddr = "0.0.0.0:8080"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Path[1:]

	if _, err := os.Stat(f); err != nil {
		f = "index.html"
	}
	http.ServeFile(w, r, f)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	fmt.Println(http.ListenAndServe(":8080", mux))
}
