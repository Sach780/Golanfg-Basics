package main

import (
    "encoding/json"
    "fmt")

var socialsites = []byte(`{
  "users": [
    {
      "name": "Elliot",
      "type": "Reader",
      "age": 23,
      "social": {
        "facebook": "https://facebook.com",
        "twitter": "https://twitter.com"
      }
    },
    {
      "name": "Fraser",
      "type": "Author",
      "age": 17,
      "social": {
        "facebook": "https://facebook.com",
        "twitter": "https://twitter.com"
      }
    }
  ]
}`)

type Users struct {
    
    Users   []User   `json:"user"`
}

type User struct {
    
    Type    string   `json:"type,attr"`
    Name    string   `json:"name"`
    Social  Social   `json:"social"`
}

type Social struct {
    
    Facebook string   `json:"facebook"`
    Twitter  string   `json:"twitter"`
    Youtube  string   `json:"youtube"`
}

func main() {

    bytes := socialsites
	 var users Users
	json.Unmarshal(bytes, &users)
	
    for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        fmt.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }

}
