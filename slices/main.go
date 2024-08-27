package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Avacado", "Orange")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// second method for initialise the slice

	marks := make([]int, 3)

	marks[0] = 99
	marks[1] = 94
	marks[2] = 92

	marks = append(marks, 56, 59)

	fmt.Println("marks slice values", marks)

	//fmt.Println(sort.IntsAreSorted(marks))
	sort.Ints(marks)
	//fmt.Println(marks)
	fmt.Println("marks slice values", marks)

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	/*
			courses[:index] refers to the slice from the start up to
			(but not including) the element at index.


			 In this case, it includes "reactjs" and "javascript".

		courses[index+1:] refers to the slice from the element after index
		to the end. In this case, it includes "python" and "ruby".

		append(courses[:index], courses[index+1:]...) effectively concatenates
		 the two slices, leaving out "swift".


	*/

}
