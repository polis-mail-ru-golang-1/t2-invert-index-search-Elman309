package controller

import (
	"log"
	"net/http"

	"github.com/Elman309/invert-index/model"
	"github.com/Elman309/invert-index/view"
)

// Controller realizes MVC controller entity
type Controller struct {
	View  view.View
	Model model.Model
}

// New returns empty controller
func New(view view.View, model model.Model) Controller {
	return Controller{
		View:  view,
		Model: model,
	}
}

// Search handles search query
func (c Controller) Search(w http.ResponseWriter, r *http.Request) {
	var err error
	searchResults := c.Model.Search(r.URL.Query().Get("q"))
	if len(searchResults) != 0 {
		err = c.View.ResultsView(searchResults, w)
	} else {
		err = c.View.SearchView(w)
	}
	if err != nil {
		log.Println("controller error:", err)
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
