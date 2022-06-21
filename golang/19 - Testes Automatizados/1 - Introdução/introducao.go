package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("RODOVIA dos Imigrantes")
	fmt.Println(tipoEndereco)
}
