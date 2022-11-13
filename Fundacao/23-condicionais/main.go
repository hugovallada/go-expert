package main

import "fmt"

func main() {
	a := 10
	b := 2
	c := 3

	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}

	if a > b && c > a { // E
		fmt.Println("C é o maior")
	}

	if a > b || c > b { // OU
		fmt.Println("B é o menor")
	}

	switch a {
	case 1:
		fmt.Println("a")
	case 2:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}

	switch {
	case a > b:
		fmt.Println("A é maior que b")
	case c > b:
		fmt.Println("C é maior que b")
	default:
		fmt.Println("B é o maior")
	}

	if b > a {
		fmt.Println("a")
	} else if a > b {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}
}
