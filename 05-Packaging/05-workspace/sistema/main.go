package main

import (
	"fmt"

	mm "github.com/hugovallada/maths"
)

func main() {
	m := mm.NewMath(5, 2)
	fmt.Println(m.Add())
}
