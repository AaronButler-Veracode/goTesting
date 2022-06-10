package main

import "fmt"

func pointers() {
	var firstName *string = new(string)

	*firstName = "Arthur"

	fmt.Println(*firstName)

}
