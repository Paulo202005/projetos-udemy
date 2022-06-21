package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	// Enquanto que
	i := 0
	for i < 3 {
		i++
		fmt.Println("Icrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	// for init + Enquanto que
	for j := 0; j < 6; j += 2 {
		fmt.Println("Icrementando j", j)
		time.Sleep(time.Second)
	}

	// for range - iteração Array
	nomes := [3]string{"João", "Davi", "Lucas"}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// for range - iteração Slice
	nomes1 := []string{"João", "Davi", "Lucas"}
	for _, nome := range nomes1 {
		fmt.Println(nome)
	}

	// for range - iteração Strig
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// for range - iteração Map
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// for loop infinito
	for {
		fmt.Println("Executando um loop infinito")
		time.Sleep(time.Second)
	}
}
