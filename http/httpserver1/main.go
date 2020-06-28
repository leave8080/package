package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}
func url(w http.ResponseWriter, r *http.Request) {
	urlinfo := r.URL
	fmt.Fprintf(w, urlinfo.Path, urlinfo.User)
	fmt.Println(urlinfo.Path)
}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/url", url)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
