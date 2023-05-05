package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums)
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)

	sum(49, 21, 324, 624)

	nums := []int{23, 63, 7567, 6586, 345}
	sum(nums...)
}
