package main

import (
	"fmt"
	"github.com/killtw/flapper/parser"
)

func main() {
	//path := "PRED-148"
	//path := "071519_871-1pon-1080p"
	//path := "072019_01-10mu-1080p"
	//path := "062819-950-carib-1080p"
	path := "heyzo_hd_2040_full"

	fmt.Println(parser.New(path))
}
