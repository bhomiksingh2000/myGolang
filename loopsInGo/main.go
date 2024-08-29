package main

import "fmt"

func main() {
	names := []string{"Bhomik", "BSA", "MNA"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for i := range names {
	// 	fmt.Println(names[i])
	// }

	// for index, name := range names {
	// 	fmt.Println("index is %v, and value is %v\n", index, name)
	// }

	for _, name := range names {
		fmt.Println("name is %v\n", name)
	}
}
