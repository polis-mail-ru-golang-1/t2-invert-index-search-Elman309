package inverted

// Index structure : token -> document name -> occurences count
type Index map[string]map[string]int

// NewIndex default constructor for Index
func NewIndex() Index {
	return make(Index)
}

// Update updates index with token from docName
func (index Index) Update(token string, docName string) {
	_, prs := index[token]
	if prs {
		index[token][docName]++
	} else {
		index[token] = make(map[string]int)
		index[token][docName] = 1
	}
}

// ProcessQuery returns map of documents related to their rank (sum of token weights)
func (index Index) ProcessQuery(query string) map[string]int {
	queryTokens := Tokenize(query)
	result := make(map[string]int)
	for _, token := range queryTokens {
		result = merge(result, index[token])
	}
	return result
}

func merge(dest map[string]int, src map[string]int) map[string]int {
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
