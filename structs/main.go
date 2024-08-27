package main

import "fmt"

func main() {

	fmt.Println(" structs are like classes in Go")

	// there is no inheritance in go
	// therefore no parent - child relationships will be there between structs

	bhomik := User{"Bhomik", 4, 24}
	fmt.Println(bhomik)
	fmt.Printf("bhomik details are: %+v\n", bhomik)
	fmt.Printf("Name is %v and age is %v.", bhomik.Name, bhomik.Age)
}

type User struct {
	Name string
	id   int
	Age  int
}
