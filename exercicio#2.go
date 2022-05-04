/*
Cada novo termo na sequência de Fibonacci é gerado pela adição dos dois termos anteriores. Começando com 1 e 2, os 10 primeiros termos serão:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, …
Considerando os termos da sequência de Fibonacci cujos valores não excedem quatro milhões, encontre a soma dos termos de valor par.*/

package main

import (
	"fmt"
)

//var n int
func main() {
	fibonacci(4000000)
}
func fibonacci(termo uint) {
	//declaração de variáveis
	var i uint
	var a uint = 1
	var b uint = 0
	var c uint = 0
	var sum uint = 0
	//lógica
	if (termo == 0) || (termo < 0) {
		fmt.Println(0)
	} else if termo == 1 {
		fmt.Println(1)
	} else {
		for i = 1; i < termo; i++ {
			c = b + a
			b = a
			a = c
			//fmt.Println(a)
			if a%2 == 0 && sum <= termo {
				sum += a
			}
		}
		//resultado
		fmt.Printf("A soma dos termos de valor par é: %v\n", sum)
	}
}
