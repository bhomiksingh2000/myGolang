package main

import "fmt"

func main() {
	myNumber := 23

	ptr := &myNumber

	fmt.Println("value of ptr", *ptr)
	fmt.Println("address of myNumber , ie the value of ptr", ptr)

	*ptr = *ptr * 2
	fmt.Println("value of ptr", myNumber)
}
