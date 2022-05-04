/*
2520 é o menor número que pode ser dividido por cada um dos números de 1 a 10 sem deixar resto.
Qual é o menor número positivo que é divisível por todos os números de 1 a 20?
*/

package main

import "fmt"

func main() {
	fmt.Println(lowestDiv())
}

func lowestDiv() int {
	c := 1
	resp := 0
	for resp == 0 {

		if c%11 == 0 && c%13 == 0 && c%14 == 0 && c%15 == 0 && c%16 == 0 && c%17 == 0 && c%18 == 0 && c%19 == 0 && c%20 == 0 {
			resp = c
		}
		c++
	}
	return resp
}
