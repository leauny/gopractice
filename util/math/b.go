package maths

import "fmt"

func sub(a, b int) int {
	return a-b
}

func init() {
	fmt.Println("Init math package 1")
}

func init() {
	fmt.Println("Init math package 2")
}
