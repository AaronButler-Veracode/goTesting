package main

import "fmt"

func arrays() {
	// first element is zero

	//long form array
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[2])

	// short form
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(arr2[2])

}
