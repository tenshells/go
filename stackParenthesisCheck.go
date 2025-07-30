/**
 * @input A : String
 * 
 * @Output Integer
 */
package main

import (
    "fmt"
)

type Node struct {
    value rune
    next *Node
}

type Stack struct {
    head *Node
}

func (s *Stack) push(new_val rune) {
    new_node := &Node{value: new_val, next: s.head}
    s.head = new_node
}

func (s *Stack) peek() rune {
    if s.head == nil {
        return 0
    } else {
        return s.head.value
    }
}

func (s *Stack) pop() {
    if s.head != nil     {
        s.head = s.head.next
    }
}

func solve(A string )  (int) {
    var s Stack
    correctOpen := map[rune]rune{'}':'{', ']':'[', ')':'('}
    for _, char := range(A) {
        if char == '{' || char == '(' || char == '[' {
            // fmt.Printf("got a char %v \t", char)
            s.push(char)
        } else if char == '}' || char == ']' || char == ')' {
            if correctOpen[char] == s.peek() {
                s.pop()
            } else {
                return 0
            }
        } 
    }
    if s.peek() == 0 {
        return 1
    }
    return 0
}

func checkValidParenthesis(A string) {
	if solve(A) == 1 {
		fmt.Printf("string %s is a valid parenthesis \n", A)
	} else {
		fmt.Printf("string %s is not a valid parenthesis \n", A)
	}
}

func main() {
	checkValidParenthesis("{[()]}")
	checkValidParenthesis("((((()))))[][]{{}}")
	checkValidParenthesis("([)]")
}
