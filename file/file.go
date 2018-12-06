package main

import "fmt"

type output string

func (o output) Client() {
	fmt.Println("Files")
}

// Backend exported
var Backend output
