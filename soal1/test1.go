package main

import "fmt"

func main() {

	var n, a int = 0, 0

	fmt.Print("Type input: ")
	fmt.Scanln(&n)

	if n >= 1 {
		for i := 0; i < n; i++ {
			a = (i * (i + 1) / 2) + 1
			fmt.Print(a, " ")
		}
	} else {
		fmt.Println("Wrong")
	}
}
