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
func New(pg *pg.DB, index Index) Model {
	InitTokenize()
	return Model{
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
		panic(err)
	}
	log.Println("Added word:", word)
	return &w
}

// AddFile adds file to DB
func (m Model) AddFile(name string) *File {
	f := File{Name: name}
	if _, err := m.PG.Model(&f).Where("name = ?", name).SelectOrInsert(); err != nil {
		panic(err)
	}
	log.Println("Added file:", name)
	return &f
}

// Upload inserts whole index to DB
func (m Model) Upload() {
	for word, files := range m.Index {
		m.Words[word] = m.AddWord(word).ID
		for file, count := range files {
			o := Occurence{
				WordID: m.Words[word],
				FileID: m.Files[file],
				Count:  count,
			}
			if _, err := m.PG.Model(&o).Insert(); err != nil {
				log.Println("tried to add existing occurence")
			} else {
				log.Println("Added occurence", word, "to", file, "count", count)
			}
		}
	}
	log.Println("Done adding occurences")
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
