package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) string {
	d[word] = definition
	return definition
}

func Search(dictionary Dictionary, word string) string {
	return dictionary[word]
}

// func main() {
// 	dictionary := Dictionary{"hello": "a nice friendly greeting"}
// 	fmt.Println(dictionary.Search("hello"))
// 	fmt.Println(dictionary.Search("goodbye"))
// }
