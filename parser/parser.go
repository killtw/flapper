package parser

import (
	"log"
	"regexp"
)

var (
	jav      = regexp.MustCompile(`[A-Za-z]{2,5}\-?\d{2,5}`)
	onepondo = regexp.MustCompile(`(\d{6}[-_]\d{3})[-_]1pon`)
	musume   = regexp.MustCompile(`(\d{6}[-_]\d{2})[-_]10mu`)
	carib    = regexp.MustCompile(`(\d{6}[-_]\d{3})[-_]carib`)
	heyzo    = regexp.MustCompile(`heyzo([-_]hd)?[-_](\d{3,4})`)
)

func New(path string) (parser string) {
	switch {
	case onepondo.MatchString(path):
		parser = "1pon"
	case musume.MatchString(path):
		parser = "10musume"
	case carib.MatchString(path):
		parser = "carib"
	case heyzo.MatchString(path):
		parser = "heyzo"
	case jav.MatchString(path):
		parser = "jav"
	default:
		log.Fatal("Nothing matched")
	}

	return
}
