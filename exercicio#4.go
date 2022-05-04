/*
Um número palindrômico lê da mesma forma nos dois sentidos. O maior palíndromo feito do produto de dois números de 2 dígitos é 9009 = 91 × 99.
Encontre o maior palíndromo feito do produto de dois números de 3 dígitos
*/
package main

import "fmt"

var biggest compare
var sl1 [][]int
var lowest compare

type compare []int

func main() {
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			x := j * i       // multiplicação entre todos o números de 3 dígitos
			z := splitInt(x) // split int
			for p := 0; p < len(z)-1; p++ {
				if (z[2] == z[3]) && (z[1] == z[len(z)-2]) && (z[0] == z[len(z)-1]) { //verificação se é palíndromo
					sl1 = append(sl1, z) // slice de palíndromos
					lowest := z

					if biggest < lowest {
						biggest = lowest
					}

				}

			}
		}

		//ordenando

		if len(sl1) == 6 {
			for a := 0; a < len(sl1)-1; a++ {
				for b := 0; b < len(sl1)-1; b++ {
					for c := 0; c < len(sl1)-1; c++ {
						for d := 0; d < len(sl1)-1; d++ {
							if sl1[a][0] < sl1[c][0] && sl1[a][5] < sl1[c][5] {
								aux := sl1[a][b]
								sl1[a][b] = sl1[c][d]
								sl1[c][d] = aux
							}
						}
					}
				}
			}
		}
	}

	if len(sl1) == 6 {
		for a := 0; a < len(sl1)-1; a++ {
			for b := 0; b < len(sl1)-1; b++ {
				if sl1[a] < sl1[b] {
					aux := sl1[a]
					sl1[a] = sl1[b]
					sl1[b] = aux
				}
			}

		}

	}

	for i, v := range sl1 {
		fmt.Println(i, v)
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
