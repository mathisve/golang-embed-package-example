package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static
var f embed.FS

func main() {

	fs := http.FileServer(http.FS(f))
	http.Handle("/", fs)

	log.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
