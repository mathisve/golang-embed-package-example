package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed folder
var f embed.FS

func main() {
	dir, err := f.ReadDir("folder")
	if err != nil {
		log.Println(err)
	}

	for _, item := range dir {
		if item.IsDir() {
			continue
		}

		info, _ := item.Info()

		i, _ := f.Open("folder/" + item.Name())
		b := make([]byte, info.Size())
		i.Read(b)

		fmt.Println(string(b))
	}
}
