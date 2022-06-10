package main

import "fmt"

func pointers2() {
	firstName := "Arthur"
	fmt.Println(firstName)

	// addressof operator
	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)
}
