/*
Um número palindrômico lê da mesma forma nos dois sentidos. O maior palíndromo feito do produto de dois números de 2 dígitos é 9009 = 91 × 99.
Encontre o maior palíndromo feito do produto de dois números de 3 dígitos
*/
package main

import (
	"fmt"
)

//var sl []int
var sl1 [][]int
var x int = 0
var z int

//var a int = 9999

func main() {

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			//fmt.Printf("%v * %v = %v\n", j, i, (j * i)) // multiplicação entre todos o números de 3 dígitos
			x = j * i
			z := splitInt(x)
			//z = strconv.Itoa(x)       // conversão de int em string
			//f := strings.Split(z, "") // conversão de string em slice of strings

			for p := 0; p < len(z)-1; p++ {

				if (z[2] == z[3]) && (z[1] == z[len(z)-2]) && (z[0] == z[len(z)-1]) { //verificação se é palíndromo
					sl1 = append(sl1, z)
				}

			}
		}

	}

	for _, v := range sl1 {

		fmt.Println(v)

	}
}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10

		/*for i, v := range slc {
			fmt.Println(i, v)
		}*/
	}
	return slc
}
