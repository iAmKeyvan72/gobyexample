package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["hello"] = 2
	m["chetori"] = 4

	fmt.Println((m))

	delete(m, "hello")
	fmt.Println((m))

	_, prs := m["hello"]
	fmt.Println(prs)

	my := map[string]string{"name": "Keyvan", "deh": "Karkabood"}
	fmt.Println(my)
}
