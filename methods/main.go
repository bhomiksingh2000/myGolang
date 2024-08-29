package main

import "fmt"

func main() {

	fmt.Println(" structs are like classes in Go")

	// there is no inheritance in go
	// therefore no parent - child relationships will be there between structs

	bhomik := User{"Bhomik", 4, 24, "b@gmail.com", true}
	// fmt.Println(bhomik)
	// fmt.Printf("bhomik details are: %+v\n", bhomik)
	fmt.Printf("Name is %v and Email is %v.", bhomik.Name, bhomik.Email)

	bhomik.getStatus()
	bhomik.setEmail()

	fmt.Printf("Name is %v and Email is %v.", bhomik.Name, bhomik.Email)
}

type User struct {
	Name   string
	id     int
	Age    int
	Email  string
	Status bool
}

func (u User) getStatus() {
	fmt.Println("User Status is :", u.Status)
}

func (u User) setEmail() {
	u.Email = "b2@gmail.com"
	fmt.Println("New Email is : ", u.Email)
}
