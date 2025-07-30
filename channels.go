package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	for i:=0; i<5; i++ {
		go func(id int) {
			ch <- fmt.Sprintf("channel has been added string from %v iteration", id)
		}(i)
	}

	for i:=0; i<5; i++ {
		msg := <-ch
		fmt.Println(msg)
	}
}
