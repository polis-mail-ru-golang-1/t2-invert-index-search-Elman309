package inverted

// Documents type define just for readability
type Documents map[string]int

// Index structure : token -> document name -> occurrences count
type Index map[string]Documents

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
			index[token] = make(Documents)
			index[token][docName] = 1
		}
	}
}

// ProcessQuery returns map of documents related to their rank (sum of token weights)
func (index Index) ProcessQuery(query string) Documents {
	queryTokens := Tokenize(query)
	result := make(Documents)
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

func (docs Documents) merge(src Documents) Documents {
	for key := range src {
		_, prs := docs[key]
		if prs {
			docs[key] += src[key]
		} else {
			docs[key] = src[key]
		}
	}
	return docs
}
