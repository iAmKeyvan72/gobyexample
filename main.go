package main

import "fmt"

func zeroVal(n int) {
	n = 0
}

func zeroPtr(n *int) {
	*n = 0
}

func main() {
	i := 10

	fmt.Println(i)

	zeroVal(i)
	fmt.Println(i)

	zeroPtr(&i)
	fmt.Println(i)
	fmt.Println(&i)

}
