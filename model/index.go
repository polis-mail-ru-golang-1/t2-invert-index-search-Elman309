package model

// Files type define just for readability
type Files map[string]int

// Index structure : token -> file name -> occurrences count
type Index map[string]Files

// NewIndex default constructor for Index
func NewIndex() Index {
	return make(Index)
}

// Update updates index with tokens from docName
func (index Index) Update(docName string, tokens ...string) {
	for _, token := range tokens {
		_, prs := index[token]
		if prs {
			index[token][docName]++
		} else {
			index[token] = make(Files)
			index[token][docName] = 1
		}
	}
}

// ProcessQuery returns map of Files related to their rank (sum of token weights)
// TODO: Planned to make it deprecated
func (index Index) ProcessQuery(query string) Files {
	queryTokens := Tokenize(query)
	result := make(Files)
	for _, token := range queryTokens {
		result = result.merge(index[token])
	}
	return result
}

// Merge merges two inverted indices
func (index Index) Merge(src Index) {
	for token := range src {
		_, prs := index[token]
		if prs {
			index[token].merge(src[token])
		} else {
			index[token] = src[token]
		}
	}
}

func (files Files) merge(src Files) Files {
	for key := range src {
		_, prs := files[key]
		if prs {
			files[key] += src[key]
		} else {
			files[key] = src[key]
		}
	}
	return files
}
