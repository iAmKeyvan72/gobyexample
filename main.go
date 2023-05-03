package main

import "fmt"

func returnTwo() (string, int) {
	return "hoy", 3
}

func main() {
	fmt.Println(returnTwo())

	_, second := returnTwo()
	fmt.Println(second)
}
