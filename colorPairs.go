package main

import "fmt"

func main() {
	var ar = []int{1, 2, 1, 2, 1, 3, 2}
	fmt.Println(sockMerchant(7, ar))
}

func sockMerchant(n int, arr []int) int {
	var cont int
	var aux int
	/*for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if arr[i] == arr[j] {
				cont += 1
				break
			}
		}
	}
	*/

	for i := 0; i < n; i++ {
		aux = arr[0]
		if aux == arr[i] {
			cont += 1
		}
	}

	return cont
}
