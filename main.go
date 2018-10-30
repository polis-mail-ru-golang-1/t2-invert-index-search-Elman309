package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sync"

	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-Elman309/inverted"
)

type rankedDocument struct {
	name string
	rank int
}

func updateFromFile(index inverted.Index, fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	tokens := inverted.Tokenize(string(data))
	for _, token := range tokens {
		index.Update(token, fileName)
	}
}

func formatRanked(rankedDocs map[string]int) {
	result := make([]rankedDocument, 0)
	for doc := range rankedDocs {
		result = append(result, rankedDocument{name: doc, rank: rankedDocs[doc]})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].rank > result[j].rank
	})
	for _, doc := range result {
		fmt.Printf(" - %s;\t совпадений - %d \n", doc.name, doc.rank)
	}
}

func testFiles(query string, fileNames ...string) {
	// keeping separate index for every file
	// merging them into single one seems to be useless (not sure)
	indices := make(map[string]inverted.Index, len(fileNames))

	// concurrently indexate files
	var wg sync.WaitGroup
	wg.Add(len(fileNames))
	for _, fileName := range fileNames {
		indices[fileName] = inverted.NewIndex()
		go func(idx inverted.Index, fn string) {
			defer wg.Done()
			updateFromFile(idx, fn)
		}(indices[fileName], fileName)
	}
	wg.Wait()
	// finished indexing

	rankedDocs := make(map[string]int, len(fileNames))
	for _, idx := range indices {
		inverted.ResultMerge(rankedDocs, idx.ProcessQuery(query))
	}
	formatRanked(rankedDocs)
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command-line arguments")
	}

	fileNames := os.Args[1 : len(os.Args)-1]
	query := os.Args[len(os.Args)-1]
	// precompiling regexp
	inverted.InitTokenize()
	testFiles(query, fileNames...)
}
