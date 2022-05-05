/*
Na sua linguagem favorita faça um algoritmo que dado uma string com letras em um mesmo
case me retorna os buracos de caracteres faltantes de A-Z

Exemplo:
"abcdf" => "e"
"acdf" => "be"
"OQRS" => "P"
"acdb" => ""
"abcz" =>  "defghijklmnopqrstuvwxy"
*/

package main

import (
	"fmt"
)

/*

func LetrasFaltantes(sequence string) {
	num := []byte(sequence)
	for i := 0; i < len(num); i++ {
		fmt.Printf("%#U\n", num[i])

	}
}

*/

func LetrasFaltantes(sequence string) []byte {
	x := []byte(sequence)
	fmt.Println(x)
	var z []byte

	for i := 0; i < len(x); i++ {
		if x[i+1]-x[i] > 1 {
			z = x[:i]

			z = append(z, x[i]+1)

		}

	}
	return z
}
func main() {

	LetrasFaltantes("abcf")

}
