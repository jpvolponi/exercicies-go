package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var reverso string
	var original int

	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			original = i * j
			aux := strings.Split(strconv.Itoa(original), "")
			//aux := splitInt(original)
			for k := len(aux) - 1; k > 0; k-- {
				reverso += aux[k]
				fmt.Println("Debug:", original, ":", reverso)
				if reverso == aux {
					fmt.Println("Debug:", reverso)
				}
			}
		}
	}

}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10

		/*
			for i, v := range slc {
				fmt.Println(i, v)
			}
		*/
	}
	return slc
}
