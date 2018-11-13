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
func (index Index) Update(fileName string, tokens ...string) {
	for _, token := range tokens {
		_, prs := index[token]
		if prs {
			index[token][fileName]++
		} else {
			index[token] = make(Files)
			index[token][fileName] = 1
		}
	}
}

// UpdateFromString ...
func (index Index) UpdateFromString(fileName string, str string) {
	tokens := Tokenize(str)
	index.Update(fileName, tokens...)
}

// IndexMerge merges two inverted indices
/* func (index Index) IndexMerge(src Index) {
	for token := range src {
		_, prs := index[token]
		if prs {
			index[token].FilesMerge(src[token])
		} else {
			index[token] = src[token]
		}
	}
} */
func IndexMerge(dest Index, src Index) {
	for token := range src {
		_, prs := dest[token]
		if prs {
			dest[token] = FilesMerge(dest[token], src[token])
		} else {
			dest[token] = src[token]
		}
	}
}

// FilesMerge ...
/* func (files Files) FilesMerge(src Files) Files {
	for key := range src {
		_, prs := files[key]
		if prs {
			files[key] += src[key]
		} else {
			files[key] = src[key]
		}
	}
	return files
} */
func FilesMerge(dest Files, src Files) Files {
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
