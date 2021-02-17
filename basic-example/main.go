package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed example.txt
var f string

//go:embed data.json
var d []byte

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	fmt.Println(f)

	var p person
	_ = json.Unmarshal(d, &p)

	fmt.Printf("%+v\n", p)
}
