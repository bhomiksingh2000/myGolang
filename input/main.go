package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// instead of try / catch , in GO we have comma / ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)

	// _: The underscore is used to ignore the second return value
	// of the ReadString method, which is an error. This is useful
	// when you're not interested in handling or checking the error
	// in this particular context.

	// therefore we can write :
	// input, error := reader.ReadString('\n')
}
