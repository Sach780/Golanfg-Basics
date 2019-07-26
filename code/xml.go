package main

import (
    "encoding/xml"
    "fmt")

var socialsites = []byte(`<users>
<user type="admin">
  <name>Elliot</name>
  <social>
	<facebook>https://facebook.com</facebook>
	<twitter>https://twitter.com</twitter>
	<youtube>https://youtube.com</youtube>
  </social>
</user>
<user type="reader">
  <name>Fraser</name>
  <social>
	<facebook>https://facebook.com</facebook>
	<twitter>https://twitter.com</twitter>
	<youtube>https://youtube.com</youtube>
  </social>
</user>
</users>`)

type Users struct {
    
    Users   []User   `xml:"user"`
}

type User struct {
    
    Type    string   `xml:"type,attr"`
    Name    string   `xml:"name"`
    Social  Social   `xml:"social"`
}

type Social struct {
    XMLName  xml.Name `xml:"social"`
    Facebook string   `xml:"facebook"`
    Twitter  string   `xml:"twitter"`
    Youtube  string   `xml:"youtube"`
}

func main() {

    bytes := socialsites
	 var users Users
	xml.Unmarshal(bytes, &users)
	
    for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        fmt.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }

}
