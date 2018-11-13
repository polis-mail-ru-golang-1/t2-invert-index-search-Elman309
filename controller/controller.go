package controller

import (
	"log"
	"net/http"

	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-Elman309/model"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-Elman309/view"
)

// Controller realizes MVC controller entity
type Controller struct {
	View  view.View
	Model *model.Model
}

// New returns empty controller
func New(view view.View, model *model.Model) Controller {
	return Controller{
		View:  view,
		Model: model,
	}
}

// SearchHandler handles search query
func (c Controller) SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	searchResults := c.Model.Search(query)
	log.Printf("requested query \"%s\", got %d result files", query, len(searchResults))
	if err := c.View.ResultsView(searchResults, query, w); err != nil {
		log.Printf("failed to create view for query %s, message: %s", query, err)
	}
}

// IndexHandler handles index page request
func (c Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := c.View.IndexView(w); err != nil {
		log.Printf("failed to create index page, message: %s", err)
	}
}

// UploadHandler ...
func (c Controller) UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		c.Model.BuildFromString(
			r.PostForm.Get("filename"),
			r.PostForm.Get("content"),
		)
		c.Model.Upload()
	}

	if err := c.View.UploadView(w); err != nil {
		log.Printf("failed to create upload page, message: %s", err)
	}
}

// AddFiles ...
func (c Controller) AddFiles(names ...string) {
	c.Model.Build(names...)
	c.Model.Upload()
}

// AddFilesAll ...
func (c Controller) AddFilesAll() {
	c.Model.BuildAll()
	c.Model.Upload()
}

/* // AddFileFromData ...
func (c Controller) AddFileFromData(fileName string, data []byte) {

} */
