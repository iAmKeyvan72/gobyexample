package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	sum := plus(45, 2342)

	fmt.Println(sum)
}
