package main

import (
	"fmt"
)

func main() {
	m := maths.NewAdd(1, 2)
	fmt.Println(m.Add())
}
