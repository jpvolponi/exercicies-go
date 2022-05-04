package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(imprimeAsterisco("bienvenido", 10))
	fmt.Println(CleanupMessage("**********bienvenido**********"))
}

func imprimeAsterisco(welcomeMsg string, x int) string {
	var ast string

	for i := 0; i < x; i++ {
		ast += "*"
	}
	msg := ast + "\n" + welcomeMsg + "\n" + ast
	return msg
}

func CleanupMessage(oldMsg string) string {
	var y string
	x := strings.Split(oldMsg, "\n")[1]
	// fmt.Println("Debug:",x, len(x))
	if strings.Contains(x, "*") {
		y = strings.Split(x, "*")[1]
		//fmt.Println("Debug:",y, len(y))
	} else {
		y = x
	}
	z := strings.TrimSpace(y)
	//fmt.Println("Debug:",z, len(z))
	return z

	/*
		Código Diogo:
			func CleanupMessage(oldMsg string) string {
				x:=strings.ReplaceAll(oldMsg,"*","")
			    return strings.TrimSpace(x)
			}
	*/
}
