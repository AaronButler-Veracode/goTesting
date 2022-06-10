package main

import "fmt"

// iota are dynamically generated at compile time, is incremented each time it is called, is unique to different constant blocks
const (
	first = iota
	second
)

const (
	third = iota
	fourth
)

func constants2() {

	fmt.Println(first, second, third, fourth)

}
