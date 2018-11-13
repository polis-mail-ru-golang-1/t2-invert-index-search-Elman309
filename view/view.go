package view

import (
	"html/template"
	"io"

	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-Elman309/model"
)

// View realizes MVC view entity
type View struct {
	IndexT   *template.Template
	ResultsT *template.Template
	UploadT  *template.Template
}

// New returns View based on given template files
func New() View {
	var view View
	var err error

	view.IndexT, err = template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		panic(err)
	}

	view.ResultsT, err = template.ParseFiles("templates/layout.html", "templates/results.html")
	if err != nil {
		panic(err)
	}

	view.UploadT, err = template.ParseFiles("templates/layout.html", "templates/upload.html")
	if err != nil {
		panic(err)
	}

	return view
}

type viewResults struct {
	Results []model.Result
	Query   string
	Empty   bool
}

// IndexView ...
func (view View) IndexView(w io.Writer) error {
	return view.IndexT.ExecuteTemplate(w, "layout", nil)
}

// ResultsView ...
func (view View) ResultsView(results []model.Result, query string, w io.Writer) error {
	return view.ResultsT.ExecuteTemplate(w, "layout",
		viewResults{
			Results: results,
			Query:   query,
			Empty:   len(results) == 0,
		})
}

// UploadView ...
func (view View) UploadView(w io.Writer) error {
	return view.UploadT.ExecuteTemplate(w, "layout", nil)
}
