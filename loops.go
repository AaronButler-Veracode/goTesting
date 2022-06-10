package main

func loops() {
	/*
		// loop till condition
		var i int
		for i < 5 {
			println(i)
			i++
			if i == 4 {
				// "continue" exits out of the current loop iteration
				continue
			}

			if i == 3 {
				// "break" command exits out of the entire loop
				break
			}
			println("Continuing...")
		}
	*/

	/*
		// Loop with condition and Post clause
		for i := 0; i < 5; i++ {
			// i is only available within scope of for loop
			println(i)
		}
	*/

	/*
		// infinite loop, only exits due to if statement
		var i int
		for {
			if i == 5 {
				break
			}
			println(i)
			i++

		}
	*/

	/*
		// loops through entire slice
		slice := []int{1, 2, 3}
		for i := 0; i < len(slice); i++ {
			println(slice[i])
		}
	*/

	/*
		// manually loop over collection
		slice := []int{1, 2, 3}
		for i, v := range slice {
			println(i, v)
		}
	*/

	/*
		// loop over collection using range using key and value
		wellKnownPorts := map[string]int{"http": 80, "https": 443}
		for k, v := range wellKnownPorts {
			println(k, v)
		}
	*/

	/*
		// loop over collection and use only first variable
		wellKnownPorts := map[string]int{"http": 80, "https": 443}
		for k := range wellKnownPorts {
			println(k)
		}
	*/
	// loop over collection using 2nd variable, by using the write only variable
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for _, v := range wellKnownPorts {
		println(v)
	}
}
