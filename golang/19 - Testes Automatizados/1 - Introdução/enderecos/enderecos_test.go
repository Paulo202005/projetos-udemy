package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// Teste de Unidade em massa (TestX...)
// go test
// go test -v / Verbose detalhado
// go test --cover / percentual do pacote atendido
// go test --coverprofile cobertura.txt / Arquivo txt com a cobertura dos testes
// go tool cover --func=cobertura.txt / Ler o arquivo de cobertura
// go tool cover --html=cobertura.txt / Ler o arquivo de cobertura detalhada
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // rodar em paralelo

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada SP120", "Estrada"},
		{"RUA SÃO Paulo", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				retornoRecebido,
				cenario.enderecoInserido,
			)
		}
	}

}

func TestNovoTeste(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
