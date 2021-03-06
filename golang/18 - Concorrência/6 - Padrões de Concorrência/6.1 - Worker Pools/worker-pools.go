package main

import "fmt"

func main() {
	tarefas := make(chan int, 40)
	resultados := make(chan int, 40)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 40; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 40; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
	close(resultados)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
