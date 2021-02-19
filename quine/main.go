package main

import (
	_ "embed"
	"fmt"
)

//go:embed main.go
var s string

func main() {
	fmt.Println(s)
}