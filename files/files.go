package files

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Elman309/invert-index/inverted"
)

func valid(info os.FileInfo) bool {
	return filepath.Ext(info.Name()) == ".txt"
}

func getValidFiles() []string {
	validFiles := make([]string, 0)
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if valid(info) {
			validFiles = append(validFiles, path)
		}
		return nil
	})
	return validFiles
}

func updateFromFile(index inverted.Index, fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	tokens := inverted.Tokenize(string(data))
	index.Update(fileName, tokens...)
}

// Build concurrently builds inverted index for all files
func Build(files ...string) inverted.Index {
	index := inverted.NewIndex()
	indexReceiver := make(chan inverted.Index, len(files))

	for _, file := range files {
		go func(fn string) {
			tempIndex := inverted.NewIndex()
			updateFromFile(tempIndex, fn)
			indexReceiver <- tempIndex
		}(file)
	}

	for range files {
		index.Merge(<-indexReceiver)
	}

	return index
}

// BuildAll is calling Build based on GetValidFiles result
func BuildAll() inverted.Index {
	return Build(getValidFiles()...)
}
