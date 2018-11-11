package model

import (
	"log"
	"regexp"
	"strings"
)

var re *regexp.Regexp

// InitTokenize precompiles regexp used by Tokenize
func InitTokenize() {
	re = regexp.MustCompile("[\\p{L}\\d']+")
	log.Println("tokenizer RE initialized")
}

// Tokenize splits str by word-like substrings defined by regexp
func Tokenize(str string) []string {
	tokenPositions := re.FindAllStringIndex(str, -1)
	if len(tokenPositions) == 0 {
		return nil
	}
	tokens := make([]string, len(tokenPositions))
	for i, pos := range tokenPositions {
		tokens[i] = strings.ToLower(str[pos[0]:pos[1]])
	}
	return tokens
}
