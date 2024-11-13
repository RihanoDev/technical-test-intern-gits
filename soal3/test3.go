package main

import "fmt"

func balancedBracket(s string) string {
	stack := []rune{}

	matchBracket := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != matchBracket[char] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Println("First Input Result: ", balancedBracket("{ [ ( ) ] }"))
	fmt.Println("Second Input Result: ", balancedBracket("{ [ ( ] ) }"))
	fmt.Println("Third Input Result: ", balancedBracket("{ ( ( [ ] ) [ ] ) [ ] }"))
}
