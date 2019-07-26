package main
import ("encoding/json"
	"fmt"
)

type Book struct {
    Title string `json:"title"`
    Author Author `json:"author"`
}
type Author struct {
	Name 		string 	`json:"Name"`
	Age 		int16 	`json:"Age"`
	Developer 	bool 	`json:"Developer"`
}

// an instance of our Book struct

func main() {
	author := Author{Name: "Sachin Mishra", Age: 25, Developer: true}
	book := Book{Title: "Learning Concurreny in Python", Author: author}

	byteArray, err := json.MarshalIndent(book, "", "  ")
if err != nil {
    fmt.Println(err)
}
fmt.Println(string(byteArray))
}