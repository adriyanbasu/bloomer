package main

import (
	"unicode"
)

type tokenKind uint

const (
	syntaxtoken tokenKind = iota
	identiferToken
	integerToken
)

type token struct {
	value string
	kind  tokenkind
}

// *(+ 13 2)  *(15 and 2 = 17) **!(identiferToken)
//

func eatwhitespace(source []rune, curser int) int {
	for curser < len(source) {
		if unicode.IsSpace(source[curser]) {
			curser++
			continue
		}
		break
	}
}

func iswhitespace(r rune) bool {

}

func lex(source []rune) []token {
	var tokens []token

	curser := 0
	for {
		curser = eatwhitespace(source, curser)
	}
}
