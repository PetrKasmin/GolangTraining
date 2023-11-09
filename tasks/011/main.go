package main

import (
	"log"
	"strings"
	"unicode"
)

func main() {
	a := "mfgah134517095aldrfgvh8h"
	a = strings.TrimFunc(a, func(r rune) bool {
		return !unicode.IsNumber(r)
	})

	log.Println(a)
}
