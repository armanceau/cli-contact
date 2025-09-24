package main

import (
	"github.com/armanceau/mini-crm/internal/app"
	"github.com/armanceau/mini-crm/internal/storage"
)

func main() {
	var store storage.Storer = storage.NewJsonStore("./")
	app.Run(store)
}
