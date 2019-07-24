package parser

import (
	"fmt"
	"strings"
)

type Carib struct {
	file File
	actor string
}

func (parser *Carib) parse() Parser {
	serial := strings.ReplaceAll(carib.FindStringSubmatch(parser.file.Name)[1], "_", "-")

	doc := GetDocumentFromUrl(fmt.Sprintf("https://www.caribbeancom.com/moviepages/%s/index.html", serial), "euc-jp")
	parser.actor = doc.Find("#moviepages > div > div:nth-child(1) > div.movie-info.section.divider > ul > li:nth-child(1) > span.spec-content > a").Text()

	return parser
}

func (parser *Carib) Go() error {
	parser.parse()

	if err := parser.file.Move(parser.actor, ""); err != nil {
		return err
	}

	return nil
}
