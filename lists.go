package main

import (
	"fmt"
)

func main() {
	intList := []int{1, 2, 3, 4, 5}
	fmt.Println(intList)
	
	stringList := []string{"hey now", "say now", "run up on my end"}
	fmt.Println(stringList)

	var dynamicList []int
	dynamicList = make([]int, 3)
	fmt.Println(dynamicList)

	dynamicList = append(dynamicList, 4)
	fmt.Println(dynamicList)
}
