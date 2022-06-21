package main

import (
	"errors"
	"fmt"
)

func main() {
	// Numeros inteiros - numero indica os bits
	// int, int8, int16, int32, int64
	// apenas int -- corresponde ao bits de seu computador
	var numero int16 = 100
	var numeroN int16 = -100
	fmt.Println(numero, numeroN)

	// Numeros inteiros não negativos - unsygned int
	// uint, uint8, uint16, uint32, uint64
	var numero2 uint16 = 100
	fmt.Println(numero2)

	// alias
	// int32 = rune
	// uint8 = byte
	var numero3 rune = 100
	var numero4 byte = 100
	fmt.Println(numero3, numero4)

	// Numeros reais (Flutuantes)
	// float32, float64
	var numeroreal float32 = 100.25
	fmt.Println(numeroreal)

	// Strings
	// string
	var str string = "Teste"
	fmt.Println(str)

	// Char
	// Precisa estar em '' e converte para ASCII (int)
	char := 'B'
	fmt.Println(char)

	// Valor inicial = valor zero
	// numeros = 0
	// string = branco
	// boll = false
	// error = <nil>

	// Numero booleano
	// bool (True, False)
	var booleano1 bool = true
	fmt.Println(booleano1)

	// Tipo Erro
	// error
	var erro error
	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro, erro1)
}
