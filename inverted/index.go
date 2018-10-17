package inverted

import (
	"io/ioutil"
)

type IndexedDocument struct {
	Name  string
	Index map[string]int
}

func NewIndexedDocument() IndexedDocument {
	var doc IndexedDocument
	doc.Name = ""
	doc.Index = make(map[string]int)
	return doc
}

func (doc IndexedDocument) Update(token string) {
	_, prs := doc.Index[token]
	if prs {
		doc.Index[token]++
	} else {
		doc.Index[token] = 1
	}
}

func (doc *IndexedDocument) UpdateFromFile(fileName string) {
	doc.Name = fileName
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	tokens := Tokenize(string(data))
	for _, token := range tokens {
		doc.Update(token)
	}
}

func (doc IndexedDocument) GetOccurencesCount(token string) int {
	return doc.Index[token]
}
