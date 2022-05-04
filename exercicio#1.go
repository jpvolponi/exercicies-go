/*
“Se listarmos todos os números naturais abaixo de 10 que são múltiplos de 3 ou 5, obtemos 3, 5, 6 e 9. A soma desses múltiplos é 23.
Encontre a soma de todos os múltiplos de 3 ou 5 abaixo de 1000.”
*/

package main

import (
	"fmt"
)

var cont int

func main() {

	for i := 0; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			cont += i

		}

	}
	fmt.Println(cont)

}
