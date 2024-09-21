package main

import (
	"github.com/MartinToruan/explore-go-in-action/9_testing-and-benchmarking/4_testing-endpoints/handler"
	"log"
	"net/http"
)

func main() {
	handler.Routes()

	log.Println("listener: Started : Listening on :4000")

	http.ListenAndServe(":4000", nil)
}
