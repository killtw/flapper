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
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.AddCookie(&http.Cookie{Name: "over18", Value: "18"})

	var client = &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", response.StatusCode, response.Status)
	}

	return response
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
