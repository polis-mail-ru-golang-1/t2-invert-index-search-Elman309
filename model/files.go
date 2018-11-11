package model

import (
	"io/ioutil"
	"os"
	"path/filepath"
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

func updateFromFile(index Index, fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	tokens := Tokenize(string(data))
	index.Update(fileName, tokens...)
}

// Build concurrently builds inverted index for all files
func (m Model) Build(files ...string) {
	//index := NewIndex()
	indexReceiver := make(chan Index, len(files))
	for _, file := range files {
		m.Files[file] = m.AddFile(file).ID
		go func(fn string) {
			tempIndex := NewIndex()
			updateFromFile(tempIndex, fn)
			indexReceiver <- tempIndex
		}(file)
	}

	for range files {
		m.Index.Merge(<-indexReceiver)
	}
}

// BuildAll is calling Build based on GetValidFiles result
func (m Model) BuildAll() {
	m.Build(getValidFiles()...)
}
