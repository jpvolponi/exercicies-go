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
var z []int
var maior int
var menor int
var resp []int

//var a int = 9999

func main() {

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			x = j * i        // multiplicação entre todos o números de 3 dígitos
			z := splitInt(x) // split int
			//fmt.Printf("%T", z)
			if (z[2] == z[3]) && (z[1] == z[len(z)-2]) && (z[0] == z[len(z)-1]) { //verificação se é palíndromo
				//sl1 = append(sl1, z) // slice de palíndromos
				maior = z[2] + z[3] + z[1] + z[len(z)-2] + z[0] + z[len(z)-1]
				if menor < maior {
					menor = maior
					resp = z
				}

			}
			/*for p := 0; p < len(z)-1; p++ {



			}
			*/
		}

	}

	/*
		sort.Slice(sl1, func(i, j int) bool {
			// edge cases
			if len(sl1[i]) == 0 && len(sl1[j]) == 0 {
				return false // two empty slices - so one is not less than other i.e. false
			}
			if len(sl1[i]) == 0 || len(sl1[j]) == 0 {
				return len(sl1[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
			}

			// both slices len() > 0, so can test this now:
			//	return sl1[i][:] < sl1[j][:]
			return ((sl1[i][0] < sl1[j][0]) && (sl1[i][1] < sl1[j][1]) && (sl1[i][2] < sl1[j][2]) && (sl1[i][3] < sl1[j][3])) //(sl1[i][len(sl1)-2] < sl1[j][len(sl1)-2]) && (sl1[i][len(sl1)-1] < sl1[j][len(sl1)-1])
		})
	*/

	/*for h := 0; h < len(sl1)-1; h++ {
		for k := 0; k < len(sl1)-1; k++ {
			if sl1[h][0] < sl1[k][0] {
				fmt.Println(sl1)
			}
		}

	}
	*/
	fmt.Println(resp)

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
