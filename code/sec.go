package main

import ("fmt"
		"net/http")

func index_handler(w http.responsewriter, r *http.request){
	fmt.Fprintf("hello guys its sachin here")
}
func main(){
http.handlerfunc("/", index_handler)
http.listenandserve(":8000", nil)
}
