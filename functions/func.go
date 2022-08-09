package main

import "fmt"

func main() {
	i := complex(3, 4)
	fmt.Println(i)

	real, imginary := real(i), imag(i) // multiple assignment
	fmt.Println(real, imginary)
}
