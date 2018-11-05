package main

import (
	"flag"

	"github.com/Elman309/invert-index/files"
	"github.com/Elman309/invert-index/inverted"
	"github.com/Elman309/invert-index/server"
)

func main() {
	addressFlag := flag.String("a", "localhost:8080", "address for http server startup")

	inverted.InitTokenize()
	index := files.BuildAll()
	idxServer := server.New(index, *addressFlag)

	idxServer.Start()
}
