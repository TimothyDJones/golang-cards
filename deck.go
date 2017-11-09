package main

import (
	"fmt"
)

// Create a new type of 'deck'
// which is a "slice" of strings
type deck []string

// In this function 'd' is called a "receiver".
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
