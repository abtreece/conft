package main

import "fmt"

type output string

func (o output) Client() {
	fmt.Println("Environment")
}

// Backend exported
var Backend output
