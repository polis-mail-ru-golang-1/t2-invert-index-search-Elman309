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

// New creates new instance of IndexServer with given parameters
func New(index inverted.Index, address string) IndexServer {
	return IndexServer{
		index:   index,
		address: address,
	}
}

// Start starts HTTP server to maintain IndexServer
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

// Search processes search queries for IndexServer
func (server IndexServer) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if len(query) == 0 {
		fmt.Fprintf(w, "Usage: %s/?q=query\n", server.address)
		return
	}
	fmt.Fprintln(w, "Query: ", query)
	result := server.index.ProcessQuery(query)
	for doc, score := range result {
		fmt.Fprintln(w, doc, " -- ", score)
	}
}
