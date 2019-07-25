package parser

import (
	"log"
	"regexp"
)

var (
	jav      = regexp.MustCompile(`[A-Za-z]{2,5}-?\d{2,5}`)
	onepondo = regexp.MustCompile(`(\d{6}[-_]\d{3})[-_]1pon`)
	musume   = regexp.MustCompile(`(\d{6}[-_]\d{2})[-_]10mu`)
	carib    = regexp.MustCompile(`(\d{6}[-_]\d{3})[-_]carib`)
	heyzo    = regexp.MustCompile(`heyzo([-_]hd)?[-_](\d{3,4})`)
)

func New(path string) (parser Parser) {
	file := NewFile(path)

	switch {
	case onepondo.MatchString(path):
		parser = &Onepondo{file: file}
	case musume.MatchString(file.Name):
		parser = &Musume{file: file}
	case carib.MatchString(path):
		parser = &Carib{file: file}
	case heyzo.MatchString(file.Name):
		parser = &Heyzo{file: file}
	//case jav.MatchString(path):
	//	return "jav"
	default:
		log.Fatal("Nothing matched")
	}

	return
}

type Parser interface {
	parse() Parser
	Go() error
}
