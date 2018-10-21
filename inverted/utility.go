package inverted

import (
	"regexp"
	"strings"
)

// Tokenize splits str by word-like substrings defined by regexp
func Tokenize(str string) []string {
	re := regexp.MustCompile("[\\w]+")
	tokenPositions := re.FindAllStringIndex(str, -1)
	tokens := make([]string, len(tokenPositions))

	for i, pos := range tokenPositions {
		tokens[i] = strings.ToLower(str[pos[0]:pos[1]])
	}

	return tokens
}
