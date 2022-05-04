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
	"strings"
)

/*

func LetrasFaltantes(sequence string) {
	num := []byte(sequence)
	for i := 0; i < len(num); i++ {
		fmt.Printf("%#U\n", num[i])

	}
}

*/

func LetrasFaltantes(sequence string) {
	x := strings.Split(sequence, "")
	for i := 0; i < len(x); i++ {
		fmt.Printf("%x\n", x[i])

	}
}
func main() {

	LetrasFaltantes("abcf")

}
