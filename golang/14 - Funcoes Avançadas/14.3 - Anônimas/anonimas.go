package main

import "fmt"

func main() {
	fmt.Println("Funções Anônimas")

	func() {
		fmt.Println("Olá Mundo!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Passando parâmetro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
	}("Passando parâmetro")
	fmt.Println(retorno)
}
