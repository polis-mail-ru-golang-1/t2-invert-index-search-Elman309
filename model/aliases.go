package model

// Word ...
type Word struct {
	ID   int    `sql:"id"`
	Word string `sql:"word"`
}

// File ...
type File struct {
	ID   int    `sql:"id"`
	Name string `sql:"name"`
}

// Occurence ...
type Occurence struct {
	ID     int `sql:"id"`
	WordID int `sql:"word_id"`
	FileID int `sql:"file_id"`
	Count  int `sql:"count"`
}

// Result ...
type Result struct {
	FileName string   `sql:"file_name"`
	CountSum int      `sql:"sum"`
	Words    []string `sql:"words,array"`
}
