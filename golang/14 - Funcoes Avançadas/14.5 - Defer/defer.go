package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Executando o função alunoEstaAprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println("Clausula Defer")

	// Defer = Adiar (adia ate a ultima execução da função)
	defer funcao1()
	funcao2()
	fmt.Println(alunoEstaAprovado(1, 8))
}
