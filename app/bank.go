package main

// Customer struct {
// 	FirstName string
// 	lastName string
// 	fullName string
// }

// if you want this struct with other file in same package than surround it using type
type (
	Customer struct {
		FirstName string
		lastName  string
		fullName  string
	}
)

func getCustomer() (customer []Customer) {
	tarun := Customer{FirstName: "tarun", lastName: "vats"}
	customer = append(customer, tarun,
		Customer{FirstName: "noah", lastName: "vats"},
		Customer{FirstName: "neha", lastName: "vats"},
		Customer{FirstName: "__", lastName: "vats"},
	)
	return customer
}
