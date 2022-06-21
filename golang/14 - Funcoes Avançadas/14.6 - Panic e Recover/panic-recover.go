package main

import (
	"fmt"
)

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	fmt.Println("Executando o função alunoEstaAprovado")
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A média é exatamente 6!")
}

func main() {
	fmt.Println("Clausulas Panic e Recover")

	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}
