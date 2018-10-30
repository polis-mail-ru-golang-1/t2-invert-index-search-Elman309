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
		result = ResultMerge(result, index[token])
	}
	return result
}
