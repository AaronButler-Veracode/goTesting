package main

import "fmt"

type user struct {
	ID        int
	FirstName string
	LastName  string
}

func struts() {

	// can be defined in the function or in the package level
	// type user struct {
	// 	ID        int
	// 	FirstName string
	// 	LastName  string
	// }

	// Expanded init
	var u user
	u.ID = 42
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u, u.FirstName)

	// implicit initialization syntax
	u2 := user{ID: 1,
		FirstName: "Arther",
		LastName:  "Dent",
	}
	fmt.Println(u2)
}
