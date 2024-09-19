package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func main() {
	c := Contact{
		Name:  "Gopher",
		Title: "programmer",
	}

	data, err := json.Marshal(c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	fmt.Println(string(data))
}
