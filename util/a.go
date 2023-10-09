package util

import "fmt"

var (
	Name = "leauny"
)

func Add(a, b int) int {
	return a + b
}

func init() {
	fmt.Println("Init util package.")
}
