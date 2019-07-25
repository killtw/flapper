package parser

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/htmlindex"
	"io"
	"log"
	"net/http"
)

func GetDocumentFromUrl(url string, charset string) *goquery.Document {
	res := getResponseFromUrl(url)
	defer res.Body.Close()

	result, err := goquery.NewDocumentFromReader(decodeHTMLBody(res.Body, charset))
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func GetJsonFromUrl(url string) map[string]interface{} {
	res := getResponseFromUrl(url)
	defer res.Body.Close()

	var result map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&result)

	return result
}

func getResponseFromUrl(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res
}

func decodeHTMLBody(body io.Reader, charset string) io.Reader {
	if charset == "" || charset == "utf-8" {
		return body
	}

	if encoding, err := htmlindex.Get(charset); err == nil {
		if name, _ := htmlindex.Name(encoding); name != "utf-8" {
			body = encoding.NewDecoder().Reader(body)
		}
	}

	return body
}
