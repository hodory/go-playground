package main

import (
	"fmt"

	"github.com/hodory/go-playground/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")

	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	dictionary.Delete(baseWord)
	findWord, searchError := dictionary.Search(baseWord)
	if searchError != nil {
		fmt.Println(searchError)
	} else {
		fmt.Println(findWord)
	}
}
