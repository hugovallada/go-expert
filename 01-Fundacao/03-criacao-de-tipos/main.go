package main

import (
	"fmt"
)

type ID int

func main() {
	var id ID

	id = 1

	fmt.Println(typeOf(id))
}

func typeOf(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case ID:
		return "ID"
	default:
		return "unknown"
	}
}
