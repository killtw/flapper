package parser

import (
	"fmt"
)

type Heyzo struct {
	file File
	actor string
}

func (parser *Heyzo) parse() Parser {
	serial := heyzo.FindStringSubmatch(parser.file.Name)[2]

	doc := GetDocumentFromUrl(fmt.Sprintf("http://www.heyzo.com/moviepages/%s/index.html", serial), "")
	parser.actor = doc.Find("#movie > div.info-bg > table > tbody > tr.table-actor > td:nth-child(2) > a").Text()

	return parser
}

func (parser *Heyzo) Go() error {
	parser.parse()

	if err := parser.file.Move(parser.actor, ""); err != nil {
		return err
	}

	return nil
}
