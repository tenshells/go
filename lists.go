package main

import (
	"fmt"
)

func main() {
	// Integer List
	intList := []int{1, 2, 3, 4, 5}
	fmt.Println(intList)
	
	// string lists
	stringList := []string{"hey now", "say now", "run up on my end"}
	fmt.Println(stringList)

	// zero value in int lists is 0, defining memory for an object using make
	var dynamicList []int
	dynamicList = make([]int, 3)
	fmt.Println(dynamicList)
	
	// lists are slices, append function
	dynamicList = append(dynamicList, 4)
	fmt.Println(dynamicList)
	
	sliceOfList := dynamicList[1:4]
	fmt.Println(sliceOfList)

	l := len(sliceOfList)
	fmt.Println("length of slice is ", l)

}
