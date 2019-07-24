package parser

import (
	"fmt"
	"strings"
)

type Musume struct {
	file File
	actor string
}

func (parser *Musume) parse() Parser {
	serial := strings.ReplaceAll(musume.FindStringSubmatch(parser.file.Name)[1], "-", "_")

	doc := GetDocumentFromUrl(fmt.Sprintf("https://www.10musume.com/moviepages/%s/index.html?pncid=1", serial), "euc-jp")
	parser.actor = doc.Find("#detail > div.detail-info > div:nth-child(1) > div.detail-info__meta > dl > dd:nth-child(8) > a").Text()

	return parser
}

func (parser *Musume) Go() error {
	parser.parse()

	if err := parser.file.Move(parser.actor, ""); err != nil {
		return err
	}

	return nil
}
