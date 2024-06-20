package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    var p fastjson.Parser
    jsonData := `{"user": {"name": "John Doe", "age": 30}}`

    value, err := p.Parse(jsonData)
    if err != nil {
        panic(err)
    }

    p1 := value.Get("user")
    fmt.Printf("User name %s\n", p1.Get("name"))
     fmt.Printf("User age %s\n", p1.Get("age"))
    userJSON := value.Get("user").String()

    user :=User{}
    if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
        panic(err)
    }
    fmt.Println(user.Name, user.Age)
    
}