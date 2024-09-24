package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal("unable to download the content: ", err)
		return
	}
	defer r.Body.Close()

	f, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal("unable to create output file: ", err)
		return
	}

	w := io.MultiWriter(f, os.Stdout)

	_, err = io.Copy(w, r.Body)
	if err != nil {
		log.Fatal("unable to print out the data: ", err)
		return
	}
}
