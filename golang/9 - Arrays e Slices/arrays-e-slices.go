package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Arrays - Lista de valores
	var array1 [5]string
	array1[0] = "Posição-1"
	fmt.Println(array1)

	array2 := [5]string{"Posição-1", "Posição-2", "Posição-3", "Posição-4", "Posição-5"}
	fmt.Println(array2)
	fmt.Println(array2[3])

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slices - Se baseia na estrutura do Array, mas tem algumas particularidades
	// Sem limite de tamanho
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2) // o Slice tb é um ponteiro

	// Arrays Internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho -- 10
	fmt.Println(cap(slice3)) // capacidade -- 11

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho -- 12
	fmt.Println(cap(slice3)) // capacidade -- 24

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // tamanho -- 5
	fmt.Println(cap(slice4)) // capacidade -- 5

	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // tamanho -- 6
	fmt.Println(cap(slice4)) // capacidade -- 12
}
