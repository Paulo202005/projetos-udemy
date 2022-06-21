package main

import "fmt"

func main() {
	// Operadores não podem ter tipos diferentes

	// Aritméticos
	/*
		+ = soma
		- = subtração
		* = multiplicação
		/ = divisão
		% = modulo
	*/
	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 10 * 5
	divisao := 10 / 4
	restoDaDivisao := 10 % 2
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	// Atribuição
	/*
		= 	(var str string = "String")
		:=	(str := "String")
	*/

	// Relacionais
	/*
		==	igual
		>		maior
		>=	maior igual
		<		menor
		<=	menor igual
		!=	diferente
	*/
	fmt.Println(1 > 2)

	// Lógicos
	/*
		&&
		||
		!
	*/
	verdadeiro, falso := true, false
	fmt.Println(!verdadeiro)
	fmt.Println(verdadeiro && falso)

	// Unários
	/*
		++ (somar + 1) 						numero++
		+= (somar valor)					numero += 10
		-- (subtrair - 1)					numero--
		-= (subtrair valor)				numero -= 10
		*= (multiplicar valor )		numero *= 10
		/= (dividir valor)				numero /= 10
		%= (modulo valor)					numero %= 10	=	numero = numero % 10
	*/

	// Ternários
	/*
		não existe no Go, deverá usar If Else
	*/
}
