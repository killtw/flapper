package main

import (
	"github.com/killtw/flapper/parser"
	"log"
	"os"
)

func main() {
	if err := parser.New(os.Args[1]).Go(); err != nil {
		log.Fatal(err)
	}
}
