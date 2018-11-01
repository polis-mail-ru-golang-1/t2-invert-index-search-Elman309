package inverted

import (
	"regexp"
	"strings"
)

var re *regexp.Regexp

// InitTokenize precompiles regexp used by Tokenize
func InitTokenize() {
	re = regexp.MustCompile("[\\w]+")
}

// Tokenize splits str by word-like substrings defined by regexp
func Tokenize(str string) []string {
	//re := regexp.MustCompile("[\\w]+")
	tokenPositions := re.FindAllStringIndex(str, -1)
	tokens := make([]string, len(tokenPositions))

	for i, pos := range tokenPositions {
		tokens[i] = strings.ToLower(str[pos[0]:pos[1]])
	}

	return tokens
}

func ResultMerge(dest map[string]int, src map[string]int) map[string]int {
	for key := range src {
		_, prs := dest[key]
		if prs {
			dest[key] += src[key]
		} else {
			dest[key] = src[key]
		}
	}
	return dest
}
