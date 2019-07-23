package parser

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/htmlindex"
	"io"
	"log"
	"net/http"
)

func GetDocumentFromUrl(url string, charset string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(decodeHTMLBody(res.Body, charset))
	if err != nil {
		log.Fatal(err)
	}

	return doc
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
