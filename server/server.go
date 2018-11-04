package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Elman309/invert-index/inverted"
)

// IndexServer ...
type IndexServer struct {
	index   inverted.Index
	address string
}

// New ...
func New(index inverted.Index, address string) IndexServer {
	return IndexServer{
		index:   index,
		address: address,
	}
}

// Start ...
func (server IndexServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.Search)

	httpServer := http.Server{
		Addr:         server.address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return httpServer.ListenAndServe()
}

// Search ...
func (server IndexServer) Search(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("q")
	if len(search) == 0 {
		fmt.Fprintf(w, "Usage: %s/?q=query\n", server.address)
		return
	}
	result := server.index.ProcessQuery(search)
	for doc, score := range result {
		fmt.Fprintln(w, doc, " -- ", score)
	}
}
