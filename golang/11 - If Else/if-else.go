package main

import "fmt"

func main() {
	fmt.Println("If Else")

	numero := -10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// if init - limita apenas na estrutura de controle (if)
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Número maior que zero")
	} else if numero2 < -10 {
		fmt.Println("Número menor que -10")
	} else {
		fmt.Println("Número entre 0 a -10")
	}
}
