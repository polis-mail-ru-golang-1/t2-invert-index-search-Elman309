package main

import (
	"fmt"
	"os"
	"sort"

	"./inverted"
)

type DocumentResult struct {
	docName         string
	occurencesCount int
}

func testFiles(text string, fileNames ...string) {
	docs := make([]inverted.IndexedDocument, 0)
	for _, fileName := range fileNames {
		doc := inverted.NewIndexedDocument()
		doc.UpdateFromFile(fileName)
		docs = append(docs, doc)
	}

	tokens := inverted.Tokenize(text)
	occurences := make([]int, len(docs))

	for _, token := range tokens {
		for i, doc := range docs {
			occurences[i] += doc.GetOccurencesCount(token)
		}
	}

	docResults := make([]DocumentResult, len(docs))
	for i := range docs {
		docResults[i] = DocumentResult{docs[i].Name, occurences[i]}
	}

	sort.Slice(docResults[:], func(i, j int) bool {
		return docResults[i].occurencesCount > docResults[j].occurencesCount
	})

	for _, res := range docResults {
		if res.occurencesCount > 0 {
			fmt.Printf(" - %s;\tсовпадений -\t %d\n", res.docName, res.occurencesCount)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command-line arguments")
	}

	fileNames := os.Args[1 : len(os.Args)-1]
	text := os.Args[len(os.Args)-1]

	testFiles(text, fileNames...)
}
