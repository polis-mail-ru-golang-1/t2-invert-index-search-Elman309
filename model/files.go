package model

import (
	"io/ioutil"
	"log"
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

// Build concurrently builds inverted index for all files
func (m Model) Build(files ...string) {
	log.Println(files) // debug
	indexReceiver := make(chan Index, len(files))
	for _, file := range files {
		_, prs := m.Files[file]
		if !prs {
			m.Files[file] = m.AddFile(file).ID
		}
		go func(fn string) {
			tempIndex := NewIndex()
			data, err := ioutil.ReadFile(fn)
			if err != nil {
				panic(err)
			}
			tempIndex.UpdateFromString(fn, string(data))
			indexReceiver <- tempIndex
		}(file)
	}

	for range files {
		IndexMerge(m.Index, <-indexReceiver)
	}
}

// BuildAll is calling Build based on GetValidFiles result
func (m Model) BuildAll() {
	m.Build(getValidFiles()...)
}

// BuildFromString ...
func (m Model) BuildFromString(fileName string, str string) {
	_, prs := m.Files[fileName]
	if !prs {
		m.Files[fileName] = m.AddFile(fileName).ID
	}
	m.Index.UpdateFromString(fileName, str)
}
