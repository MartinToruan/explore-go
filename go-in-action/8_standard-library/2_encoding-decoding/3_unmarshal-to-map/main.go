package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var jsonData = `
{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

// If you don't know the structure of the JSON you're working on or you want more flexibility,
// you can use map[string]interface{}
func main() {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println("Name: ", data["name"])
	fmt.Println("Title: ", data["title"])
	fmt.Println("Contact")
	fmt.Println("Home: ", data["contact"].(map[string]interface{})["home"])
	fmt.Println("Home: ", data["contact"].(map[string]interface{})["cell"])
}
