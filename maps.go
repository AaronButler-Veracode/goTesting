package main

import "fmt"

func maps() {
	// a map that maps string (key) data types to an integer (value)
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)
}
