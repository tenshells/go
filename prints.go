package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello Go checkout-ers!\n")

	someInt := 345
	fmt.Println("One way to print number by ", someInt)

	fmt.Printf("Another way to print variables: %v\n", someInt)
	
	someString := "BROTHER"
	fmt.Println("multiple args accepted : ", someInt, " ", someString)

	fmt.Printf("diff identifiers for string and int, string %s and int/others %v\n", someString, someInt)
}
