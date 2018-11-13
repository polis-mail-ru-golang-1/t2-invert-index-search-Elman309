package model

// Word alias for public.words row
type Word struct {
	ID   int    `sql:"id"`
	Word string `sql:"word"`
}

// File alias for public.files row
type File struct {
	ID   int    `sql:"id"`
	Name string `sql:"name"`
}

// Occurence alias for public.occurences row
type Occurence struct {
	ID     int `sql:"id"`
	WordID int `sql:"word_id"`
	FileID int `sql:"file_id"`
	Count  int `sql:"count"`
}

// Result alias for search query result
type Result struct {
	FileName string   `sql:"file_name"`
	CountSum int      `sql:"sum"`
	Words    []string `sql:"words,array"`
}
