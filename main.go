package main

import (
	"github.com/Elman309/invert-index/files"
	"github.com/Elman309/invert-index/inverted"
	"github.com/Elman309/invert-index/server"
)

func main() {
	inverted.InitTokenize()
	index := files.BuildAll()
	idxServer := server.New(index, "localhost:8080")
	idxServer.Start()
}
