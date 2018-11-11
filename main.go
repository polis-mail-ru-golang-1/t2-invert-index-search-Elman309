package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Elman309/invert-index/config"
	"github.com/Elman309/invert-index/controller"
	"github.com/Elman309/invert-index/model"
	"github.com/Elman309/invert-index/view"
	"github.com/go-pg/pg"
)

func main() {
	config := config.Load()
	db := pg.Connect(&pg.Options{
		Addr:     config.Address,
		User:     config.Username,
		Password: config.Password,
		Database: config.Database,
	})
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
	s.Start(config.ServerAddress)
	s.controller.AddFilesAll()
	log.Println("Done building index")
}

// Server ...
type Server struct {
	controller controller.Controller
}

// Start ...
func (s Server) Start(address string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.controller.Search)
	mux.HandleFunc("/search/", s.controller.Search)

	server := http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}
