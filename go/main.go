package main

import (
	"fmt"
	"math"
)

const (
	N  = 31
	dx = 1
	x0 = 4
	d  = 4
)

type Poisson struct {
	u         [N][N]float64
	d, x0, dx float64
}

func NewPoisson(d, x0, dx float64) *Poisson {
	result := &Poisson{}
	return result
}

func (p *Poisson) U(i, j int) float64 {
	return p.u[i][j]
}

func (p *Poisson) Rho(i, j int) float64 {
	x, y := float64(i), float64(j)
	d := p.d
	x0 := p.x0
	return math.Exp(-(math.Pow((x-x0), 2)+y*y)/(d*d)) - math.Exp(-(math.Pow((x+x0), 2)+y*y)/(d*d))
}

func (p *Poisson) S(i, j int, delta float64) float64 {
	uij := p.U(i, j) + delta
	a1 := p.U(i-1, j) + p.U(i+1, j) + p.U(i, j-1) + p.U(i, j+1)
	a1 *= 2 * uij
	a2 := 4 * uij * uij
	a3 := p.Rho(i, j) * uij * math.Pow(p.dx, 2)

	return -0.5*(a1-a2) - a3
}

func main() {
	p := NewPoisson(d, x0, dx)
	fmt.Println(p)
}
