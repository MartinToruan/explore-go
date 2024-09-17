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
}
`

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// If you got a json in a string format,
// you can use unmarshall function to decode the json into your data structure
func main() {
	var c Contact
	err := json.Unmarshal([]byte(jsonData), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(c)
}
