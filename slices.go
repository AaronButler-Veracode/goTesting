package main

import "fmt"

func slices() {

	/*
		// slices point to an underlying array
		arr := [3]int{1, 2, 3}
		slice := arr[:]
		arr[1] = 42
		slice[2] = 27
		// changing the slice changes the underlying array
		fmt.Println(arr, slice)
	*/

	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 4, 42, 27)
	// slice resizing is managed by the underlying go lang, however you can instruct it to set a default size on initialization.
	fmt.Println(slice)

	// dynamically size or selecting sub-slices
	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)

}
