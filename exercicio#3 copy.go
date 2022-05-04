/*
Os fatores primos de 13195 são 5, 7, 13 e 29.
Qual é o maior fator primo do número 600851475143?
*/

package main

import (
	"fmt"
)

func main() {
	fatPrimo(13195)
}

func fatPrimo(num int) {

	cont := 0
	cont2 := 0
	fat := 0
	fat2 := 0
	var slice []int
	var slice2 []int
	//var teste int

	for i := 2; i < num; i++ {
		if num%i == 0 {
			cont = i
			fat = num / i
			slice = append(slice, cont)

			for j := 2; j < cont; j++ {
				if cont%j == 0 {
					fat2 = cont / j
					cont2 = j
					//fmt.Println(cont, "/", j, " = ", cont2)
					break
				}

			}

		}
	}
	slice2 = append(slice2, slice[len(slice)-1])
	fmt.Println(cont)
	fmt.Println(cont2)
	//fmt.Println(teste)
	//fmt.Println(slice)
	//fmt.Println(slice2)
	fmt.Println(fat)
	fmt.Println(fat2)

}
