/*
Os fatores primos de 13195 são 5, 7, 13 e 29.
Qual é o maior fator primo do número 600851475143?
*/

package main

import "fmt"

func main() {
	var num uint = 13195
	fatoresPrimos(num)

}

func fatoresPrimos(fimLoop uint) {
	//variáveis
	var resultado []uint
	var iniLoop uint = 1
	var quociente uint
	var divisor uint

	//lógica
	for iniLoop <= fimLoop {
		iniLoop++
		if fimLoop%iniLoop == 0 {
			quociente = fimLoop / iniLoop          //quociente
			divisor = iniLoop                      //divisor
			iniLoop = 1                            //reset do loop
			fimLoop = quociente                    //novo parâmetro pro "while"
			resultado = append(resultado, divisor) //slice de fatores
		}

	}
	//resultado
	fmt.Println(resultado)
	fmt.Println(resultado[len(resultado)-1])

}
