package main

import "fmt"

func main() {
	fmt.Println("Hello Pointer")

	var firstName *string = new(string)
	// firstName = "tarun" // error , pointer should store addresss

	*firstName = "tarun"
	fmt.Println(firstName)

	secondPointer := firstName
	fmt.Println(secondPointer, *secondPointer)
}
