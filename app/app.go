package main

import (
	"fmt"
	"time"
)

func main() {
	customer := getData(1)
	fmt.Println("Hello World " + customer)
	customer = getData(2)
	fmt.Println("Hello World " + customer)

	customer = getData(3)
	fmt.Println("Hello World " + customer)

	fmt.Println(customers())

	fmt.Println(getCustomer())

}

func getData(customerID int) (customerName string) {
	// here var is used to intialize , if you do not want to use var but intialize use := instead
	var Firstname = "Tarun"
	LastName := "Vats"
	var fullName string // varibale declaration but empty string intialization
	fullName = Firstname + " " + LastName

	if customerID == 1 {
		return fullName

	} else if customerID == 2 {
		return "Tarun chor"

	} else {
		return ""
	}

	return fullName
}

func customers() (customers []string) {
	customers = []string{"Marcel james", "sukhit cheeama"}

	customers = append(customers, "tarun vats")
	customers = append(customers, "Noah vats")
	customers = append(customers, "Neha vats")
	customers = append(customers, "--- vats")
	fmt.Println(len(customers))

	for {
		fmt.Println("loop for 1 second")
		time.Sleep(1)
		break
	}

	for x := 0; x < 10; x++ {
		fmt.Print(x)
	}

	for _, customer := range customers {
		fmt.Println(customer)
	}

	return customers
}
