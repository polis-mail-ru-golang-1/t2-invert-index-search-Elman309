package view

import (
	"html/template"
	"io"
	"log"

	"github.com/Elman309/invert-index/model"
)

// View realizes MVC view entity
type View struct {
	LayoutTemplate *template.Template
}

// New ...
func New() View {
	var view View
	var err error

	view.LayoutTemplate, err = template.ParseFiles("view/layout.html")
	if err != nil {
		panic(err)
	}

	return view
}

// Results ...
type Results struct {
	Results []model.Result
	Empty   bool
}

// ResultsView ...
func (view View) ResultsView(results []model.Result, w io.Writer) error {
	log.Println("view with results called")
	return view.LayoutTemplate.
		Execute(w, Results{Results: results, Empty: false})
}

// SearchView ...
func (view View) SearchView(w io.Writer) error {
	log.Println("empty view called")
	return view.LayoutTemplate.
		Execute(w, Results{Empty: true})
}
