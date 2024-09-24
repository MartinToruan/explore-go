package handler_test

import (
	"encoding/json"
	"fmt"
	"github.com/MartinToruan/explore-go-in-action/9_testing-and-benchmarking/4_testing-endpoints/handler"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func init() {
	handler.Routes()
}

func TestSendJSON(t *testing.T) {
	req, err := http.NewRequest("GET", "/sendjson", nil)
	if err != nil {
		log.Fatal("\tShould be able to create a request.", ballotX, err)
	}
	t.Log("\tShould be able to create a request.", checkMark)

	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, req)

	if rw.Code != 200 {
		t.Fatal("\tShould have \"200\"", ballotX, rw.Code)
	}
	t.Log("\tShould receive \"200\"", checkMark)

	u := struct {
		Name  string
		Email string
	}{}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		t.Fatal("\tShould decode the response.", ballotX)
	}
	t.Log("\tShould decode the response.", checkMark)

	if u.Name == "Kristopel" {
		t.Log("\tShould have a Name.", checkMark)
	} else {
		t.Error("\tShould have a Name.", ballotX, u.Name)
	}

	if u.Email == "kristopel@gmail.com" {
		t.Log("\tShould have an Email.", checkMark)
	} else {
		t.Error("\tShould have an Email.", ballotX, u.Email)
	}
}

// ExampleSendJSON provides a basic example
func ExampleSendJSON() {
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	rw := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rw, r)

	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	// Use fmt to write to stdout to check the output.
	fmt.Println(u)

	// Output:
	// {Kristopel kristopel@gmail.com}
}
