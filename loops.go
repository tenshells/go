package main

import (
	"fmt"
)

func main() {
	// goes uptil 5: 0 to 4
	fmt.Println("Basic running in loop: ")
	for i := range(5) {
		fmt.Println(i)
	}
	
	fmt.Printf("\n\nCountries i have visited with controlled break:\n")
	CountriesIWantTickedOffBucketList := []string{
		"India", "UAE", "Malaysia", "France", "Switzerland", "Oman", "Sri Lanka", "Singapore", "Hungary", "Netherlands"}
	
	
	for index, country := range(CountriesIWantTickedOffBucketList) {
		if country == "Hungary" {
			break
		} else {
			continue
		}

		fmt.Printf("%vth country I have visited is %s\n", index, country)
	}


	i := 0
	j := 10

	fmt.Printf("\n\nfor is only loop in go, its very flexible\n")

	for i<=j {
		fmt.Printf("%v<=%v\t", i, j)
		i++
	}

	fmt.Printf("\n\n endless loops by just skipping any params in for! ")
	i = 0
	for {
		if i <= j {
			fmt.Printf("Endless loop happening!\n")
			i++
			continue
		} else {
			fmt.Printf("Endless loop finally stopped\n")
			break
		}
	}


	
}
