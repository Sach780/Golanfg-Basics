package main

import ("fmt"
		"net/http"
		"encoding/json"
		"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article 


func AllArticles(w http.ResponseWriter, r *http.Request) {
	
	articles:=Article{Title: "Cloud Computing", Desc: "Cloud description", Content: "Services and Model"}
	articles1:=Article{Title: "Docker", Desc: "Docker description", Content: "Commands Knowledge"}
	articles2:=Article{Title: "GITHUB", Desc: "GIT and GITHUB description", Content: "Practical Knowledge"}
	articles3:=Article{Title: "Golang", Desc: "Go Lang description", Content: "Hello C2C its in a JSON format"}
	
	

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
	json.NewEncoder(w).Encode(articles1)
	json.NewEncoder(w).Encode(articles2)
	json.NewEncoder(w).Encode(articles3)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>hello guys its sachin here!</h1>
		<p> and today we study about golang </p>
		<p> golang is fast as cmp to c++</p>`)
}

func main(){

	myRouter := mux.MewRouter().StrictSlash(true)
myRouter.HandleFunc("/", index_handler)
myRouter.HandleFunc("/articles", AllArticles )
myRouter.ListenAndServe(":8000",myRouter)
}