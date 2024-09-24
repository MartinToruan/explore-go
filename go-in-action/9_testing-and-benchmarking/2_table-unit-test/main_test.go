package main

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	urls := []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.goinggo.net/feeds/posts/default?alt=rss",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content")
	{
		for _, url := range urls {
			t.Logf("When checking \"%s\" for status code \"%d\"", url.url, url.statusCode)

			resp, err := http.Get(url.url)
			if err != nil {
				t.Fatal("\t\tShould be able to Get the url.", ballotX, err)
			}
			t.Log("\t\tShould be able to Get the url", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == url.statusCode {
				t.Logf("\t\tShould have a \"%d\" status. %v", url.statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould have a \"%d\" status %v %v", url.statusCode, ballotX, resp.StatusCode)
			}

		}
	}
}
