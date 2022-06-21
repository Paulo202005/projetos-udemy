package main

import "fmt"

func main() {
	canal := make(chan string, 2) // canal com buffer (2)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	fmt.Println(mensagem)

	mensagem2 := <-canal
	fmt.Println(mensagem2)
}
