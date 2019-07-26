package main

import ("fmt"
		"log"
		"net/http"
		"encoding/json")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>hello guys its sachin here!</h1>
		<p> and today we study about golang </p>
		<p> golang is fast as cmp to c++</p>`)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "right now i'm playing for VCA!")
}

func main(){
http.HandleFunc("/", index_handler)
http.HandleFunc("/about/", about_handler)
http.ListenAndServe(":8000", nil)
}