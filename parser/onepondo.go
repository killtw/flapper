package parser

import (
	"fmt"
	"strings"
)

type Onepondo struct {
	file File
	actor string
}

func (parser *Onepondo) parse() Parser {
	serial := strings.ReplaceAll(onepondo.FindStringSubmatch(parser.file.Name)[1], "-", "_")

	result := GetJsonFromUrl(fmt.Sprintf("https://www.1pondo.tv/dyn/phpauto/movie_details/movie_id/%s.json", serial))
	parser.actor = result["Actor"].(string)

	return parser
}

func (parser *Onepondo) Go() error {
	parser.parse()

	if _, err := parser.file.Move(parser.actor, ""); err != nil {
		return err
	}

	return nil
}
