package main

import "fmt"

func divisivelPor(x int) bool {
	var isDiv bool
	for i := 2; i < x; i++ {
		if x%i == 0 {
			isDiv = true

		}
	}
	return isDiv
}

func main() {
	fmt.Println(divisivelPor(19))
}
