package main

import "fmt"

func main() {

	list := []int{5, 2, 66}
	sum := 0

	for _, num := range list {
		sum += num
	}

	fmt.Println("sum is: ", sum)

	for index, num := range list {
		if index == 2 {
			fmt.Println("this is ", index+1, "rd item which is: ", num)
		}
	}

	objs := map[string]string{"first": "hey", "second": "bro", "third": "how are you?"}
	for key, value := range objs {
		fmt.Println(key, ": ", value)
	}

	for key := range objs {
		fmt.Println(key)
	}

}
