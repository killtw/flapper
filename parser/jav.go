package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Jav struct {
	file File
	actor string
	serial string
	cover string
}

func (parser *Jav) parse() Parser {
	serial := jav.FindString(parser.file.Name)

	doc := GetDocumentFromUrl(fmt.Sprintf("http://www.m34z.com/ja/vl_searchbyid.php?keyword=%s", serial), "")

	if len(doc.Find(".video").Nodes) > 0 {
		url, _ := doc.Find(".video").First().Find("a").Attr("href")
		doc = GetDocumentFromUrl(strings.ReplaceAll(url, "./", "http://www.m34z.com/ja/"), "")
	}

	parser.serial = doc.Find("div#video_id td.text").Text()
	parser.cover, _ = doc.Find("img#video_jacket_img").Attr("src")
	var actors []string
	doc.Find("span.star a").Each(func(i int, sel *goquery.Selection) {
		actors = append(actors, sel.Text())
	})
	parser.actor = strings.Join(actors, ", ")

	if len(parser.actor) == 0 {
		parser.actor = fmt.Sprintf("[%s]", strings.ReplaceAll(doc.Find("span.label a").Text(), "/", ":"))
	}

	return parser
}

func (parser *Jav) Go() error {
	parser.parse()
	dir := filepath.Join(parser.actor, parser.serial)

	if _, err := parser.file.Move(dir, parser.serial); err != nil {
		return err
	}
	res, err := http.Get("https:" + parser.cover)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	f, err := os.Create(fmt.Sprintf("%s/%s%s", parser.file.Dir, parser.serial, filepath.Ext(parser.cover)))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = io.Copy(f, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
