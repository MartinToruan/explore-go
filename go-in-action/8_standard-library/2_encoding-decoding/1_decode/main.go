package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

// If you got a json file response from an API, you can use Decode function to decode the data into your data structure
func main() {
	url := "http://ajax.googleapis.comajax/services/search/web?v=1.0&rsz=8&q=golang"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("ERROR: ", err)
	}
	defer resp.Body.Close()

	var res gResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatalln("Error while decoding the result: ", err)
	}

	fmt.Println(res)
}
