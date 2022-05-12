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

func LetrasFaltantes(sequence string) string {
	word := []byte(sequence)
	temp := []byte(sequence)
	var result = make([]byte, len(word)+1)

	for i := 0; i < len(word)-1; i++ {

		if word[i+1]-word[i] > 1 {
			//result = word[:i+1]

			result = append(word[:i+1], word[i]+1)
			fmt.Println("debug result: ", result)
			result = append(result, temp[i+1:len(temp)]...)
			word = result
			fmt.Println("debug word: ", word)
			fmt.Println("debug result: ", result)

		}

	}
	return string(result)
}
func main() {

	fmt.Println(LetrasFaltantes("abcf"))

}
