package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-Elman309/config"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-Elman309/controller"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-Elman309/model"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-Elman309/view"
)

func main() {
	config := config.Load()

	// database connection
	db := pg.Connect(&pg.Options{
		Addr:     config.Address,
		User:     config.Username,
		Password: config.Password,
		Database: config.Database,
	})
	if db == nil {
		log.Fatalf("failed to connect to db at %s", config.Address)
	} else {
		log.Printf("successfully connected to db at %s", config.Address)
	}
	defer db.Close()

	m := model.New(
		db,
		model.NewIndex(),
	)
	c := controller.Controller{
		View:  view.New(),
		Model: m,
	}

	s := Server{controller: c}
	if config.FirstStart {
		c.AddFilesAll()
	}
	s.Start(config.ServerAddress)
}

// Server ...
type Server struct {
	controller controller.Controller
}

// Start ...
func (s Server) Start(address string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.controller.IndexHandler)
	mux.HandleFunc("/search/", s.controller.SearchHandler)
	mux.HandleFunc("/upload/", s.controller.UploadHandler)

	server := http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("starting http server at %s", address)
	return server.ListenAndServe()
}
