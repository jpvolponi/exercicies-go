/*
Os fatores primos de 13195 são 5, 7, 13 e 29.
Qual é o maior fator primo do número 600851475143?
*/

package main

func main() {
	fatPrimo(13195)
}

func fatPrimo(num int) {

	cont := 1
	//	cont2 := 0
	//var fat int = 0
	//fat2 := 0
	var slice []int
	var slice2 []int
	//var teste int

	for i := 2; i < num; i++ {
		if num%i == 0 { // 35
			//cont = i
			//fat = num / i
			slice = append(slice, i)
			//cont *= slice[i]

			//if i*(i-1) < num {
			//	slice2 = append(slice2, i)
			//}
			//fmt.Println(cont, "--") // = 5

			/*for j := 2; j < i; j++ {
				if i%j == 0 {
					fmt.Println(j)
				}
			}*/

		}

	}
	//fmt.Println(slice)
	/*for i, v := range slice {
		if slice[i]*(slice[i]-1) < num {
			cont *= slice[i]
			fmt.Println(i, v)
		}

	}*/

}
