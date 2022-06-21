package formas

import (
	"math"
)

// Forma xxxx
type Forma interface {
	Area() float64
}

// Retangulo xxxx
type Retangulo struct {
	Altura  float64
	Largura float64
}

// Area xxxx
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

// Circulo xxxx
type Circulo struct {
	Raio float64
}

// Area xxxx
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
