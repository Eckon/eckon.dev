package main

import (
	"github.com/Eckon/eckon.dev/src/server"
	"github.com/Eckon/eckon.dev/src/template"
)

func main() {
	template.Initialize()
	server.Start()
}
