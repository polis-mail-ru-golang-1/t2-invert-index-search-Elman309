package model

import (
	"log"

	"github.com/go-pg/pg"
)

// Model realizes MVC model entity
type Model struct {
	PG    *pg.DB
	Index Index
	Files map[string]int
	Words map[string]int
}

// New inits and returns empty model
func New(pg *pg.DB, index Index) *Model {
	InitTokenize()
	return &Model{
		PG:    pg,
		Index: index,
		Files: make(map[string]int),
		Words: make(map[string]int),
	}
}

// AddWord adds word to DB
func (m Model) AddWord(word string) *Word {
	w := Word{Word: word}
	if _, err := m.PG.Model(&w).Where("word = ?", word).SelectOrInsert(); err != nil {
		log.Printf("failed select or insert word \"%s\", message: %s", word, err)
	}
	log.Printf("processed word \"%s\"", word)
	return &w
}

// AddFile adds file to DB
func (m Model) AddFile(name string) *File {
	f := File{Name: name}
	_, err := m.PG.Model(&f).Where("name = ?", name).SelectOrInsert()
	if err != nil {
		log.Printf("failed select or insert file \"%s\", message: %s", name, err)
	}
	log.Printf("processed file \"%s\"", name)
	return &f
}

// GetWords downloads words from public.words table
func (m Model) GetWords() {
	var words []Word
	if err := m.PG.Model((*Word)(nil)).Select(&words); err != nil {
		panic(err)
	}
	for _, word := range words {
		m.Words[word.Word] = word.ID
	}
}

// GetFiles downloads files from public.files table
func (m Model) GetFiles() {
	var files []File
	if err := m.PG.Model((*File)(nil)).Select(&files); err != nil {
		panic(err)
	}
	for _, file := range files {
		m.Files[file.Name] = file.ID
	}
}

// Upload inserts whole index to DB
func (m *Model) Upload() {
	var occurences []Occurence
	m.GetWords()
	m.GetFiles()
	for word, files := range m.Index {
		_, prs := m.Words[word]
		if !prs {
			// issue это работает долго
			m.Words[word] = m.AddWord(word).ID
		}
		for file, count := range files {
			occurences = append(
				occurences,
				Occurence{
					WordID: m.Words[word],
					FileID: m.Files[file],
					Count:  count,
				})
		}
	}

	if err := m.PG.Insert(&occurences); err != nil {
		log.Printf("failed to insert occurences, message: %s", err)
	}
	m.Index = NewIndex()
}

// Search makes search query to DB
func (m Model) Search(query string) []Result {
	var results []Result

	words := Tokenize(query)
	if len(words) == 0 {
		return results
	}

	// main search query
	if err := m.PG.Model(&Occurence{}).
		ColumnExpr("(SELECT name FROM files WHERE id = occurence.file_id) as file_name").
		ColumnExpr("SUM(occurence.count) as sum").
		ColumnExpr("array_agg(occurence.word_id) as words").
		Join("JOIN words on occurence.word_id = words.id").
		Where("words.word in (?)", pg.In(words)).
		Group("occurence.file_id").
		Order("sum DESC").
		Select(&results); err != nil {
		log.Println(err.Error())
	}

	return results
}
