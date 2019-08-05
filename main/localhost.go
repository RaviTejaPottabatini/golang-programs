package main

import ("fmt"
"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>this is me<h1>")
}

func me_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is also me")
}
func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/hey",me_handler)
	http.ListenAndServe(":8000", nil) 
}