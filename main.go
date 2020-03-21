package main

import (
	"fmt"

	"github.com/hodory/go-playground/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
