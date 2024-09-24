package handler

import (
	"encoding/json"
	"net/http"
)

func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

func SendJSON(rw http.ResponseWriter, req *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Kristopel",
		Email: "kristopel@gmail.com",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
