package main

import (
	"github.com/killtw/flapper/config"
	"github.com/killtw/flapper/parser"
	"log"
	"os"
)

func main() {
	conf, err := config.LoadConfig("")
	if err != nil {
		log.Fatal(err)
	}
	file := parser.NewFile(os.Args[1], conf)

	if err := parser.New(file).Go(); err != nil {
		log.Fatal(err)
	}
}
