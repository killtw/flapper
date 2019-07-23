package main

import (
	"fmt"
	"github.com/killtw/flapper/parser"
)

func main() {
	doc := parser.GetDocumentFromUrl("https://www.10musume.com/moviepages/072019_01/index.html?pncid=1", "euc-jp")
	actor := doc.Find("#detail > div.detail-info > div:nth-child(1) > div.detail-info__meta > dl > dd:nth-child(8) > a").Text()
	fmt.Println(actor)
}
