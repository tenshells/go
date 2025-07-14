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

type Address2 struct {
	Name string
	HouseNumber string
	Street string
	City string
	State string
	Pincode int
	Country string
}


func (a Address2) String() string {
	return fmt.Sprintf("%s, %s\n%s\n%s, %s\n%v\n%v", a.Name, a.HouseNumber, a.Street, a.City, a.State, a.Pincode, a.Country)
}

func main() {
	fakeAddress := Address{"401", "213/B", "22nd Cross Road, HSR Layout", "Bangalore", "Karnataka", 560102, "India"}
	fmt.Println("Different ways to print types: ")
	
	fmt.Printf("yo my address is %v \n", fakeAddress)
	fmt.Printf("yo my address with var names is %+v \n", fakeAddress)
	fmt.Printf("yo my address with var, package names and types is %#v \n", fakeAddress)

	// making new type Address2 to show String() func overriding; %v and %+v will take in String() func override
	
	fakeAddress2 := Address2{"401", "213/B", "22nd Cross Road, HSR Layout", "Bangalore", "Karnataka", 560102, "India"}
	fmt.Printf("\nbetter way to print address:\n%v\n", fakeAddress2)
}

