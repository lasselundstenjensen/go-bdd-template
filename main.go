package main

import "fmt"

func main() {
	fmt.Println("Main executed")
	fmt.Println("Adding 2 to 5 is", add(2, 5))
	fmt.Println("Subtracting 7 from 6 is", subtract(6, 7))
}

func add(a, b int64) int64 {
	return a + b
}

func subtract(a, b int64) int64 {
	return a - b
}
