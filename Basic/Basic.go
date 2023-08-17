package main

import (
	"fmt"
	"time"
)

const cont = 50

func main() {
	fmt.Println("Hello World!")

	// Variables
	nameOne := "GO LANDED"
	var nameTwo string = "GO LANDED"

	fmt.Println(nameOne, nameTwo)
	fmt.Println("cont", cont)
	var age = 2007
	fmt.Println("Println", age)

	fmt.Printf("Hey %v %T\n", cont, cont)

	// DATA TYPES
	var username string = "Miguel"
	yearBorn := 1999

	fmt.Printf("Hi %v, age %v\n", username, time.Now().Year()-yearBorn)
	var name string
	fmt.Printf("Direccion of name %v\n", &name)
	_, err := fmt.Scan(&name) // indicate where Go must store the value
	if err != nil {
		return
	}
	fmt.Printf("Hi %v, age %v\n", name, time.Now().Year()-yearBorn)

}
