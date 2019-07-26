package main

import ("fmt"
		"net/http")



func main(){
http.HandleFunc("/", index_handler)
http.ListenAndServe(":8000", nil)
}
func index_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello guys its sachin here! and today we doing some calculation program")
	fmt.Println(add(20, 30))
    fmt.Printf("addition: %d\n")
}

func add(x int, y int) int {
    total := 0
    total = x + y
    return total
    
}