package main 

import (
	"fmt"
)

type Address struct {
	Name string
	HouseNumber string
	Street string
	City string
	State string
	Pincode int
	Country string
}

func main() {
	fakeAddress := Address{"401", "213/B", "22nd Cross Road, HSR Layout", "Bangalore", "Karnataka", 560102, "India"}
	fmt.Println("Different ways to print types: ")
	fmt.Printf("yo my address is %v \n", fakeAddress)
	fmt.Printf("yo my address with var names is %+v \n", fakeAddress)
	fmt.Printf("yo my address with var, package names and types is %#v \n", fakeAddress)
}

