/*
Portanto, a diferença entre a soma dos quadrados dos dez primeiros números naturais e o quadrado da soma é .
Encontre a diferença entre a soma dos quadrados dos primeiros cem números naturais e o quadrado da soma.
*/
package main

import (
	"fmt"
)

func main() {
	sumSquare(100)
}

func sumSquare(x int) {
	var sumOfSquares int
	var squareOfSum int
	var sum int
	for i := 1; i <= x; i++ {
		sumOfSquares += (i * i)
		sum += i
	}
	squareOfSum = sum * sum
	fmt.Println("\nSoma dos quadrados =", sumOfSquares)
	fmt.Println("\nQuadrado da soma =", squareOfSum)
	fmt.Println("\nDiferença entre a Soma dos quadrados e o Quadrado da som =", squareOfSum-sumOfSquares)
}
