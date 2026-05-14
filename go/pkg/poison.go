package pkg

import (
	"image"
	"math"
	"sort"
)

const (
	N = 31
)

type sdelta struct {
	s, delta float64
}

type Poisson struct {
	u                  [2*N + 1][2*N + 1]float64
	d, x0, dx          float64
	optimizationDeltas [3]float64
}

func NewPoisson(d, x0, dx float64, optimizationDeltas [3]float64) *Poisson {
	result := &Poisson{
		x0:                 x0,
		d:                  d,
		dx:                 dx,
		optimizationDeltas: optimizationDeltas,
	}
	return result
}

func (p *Poisson) GetUMap() image.Image {
	return Float64ToGrayImage(p.u)
}

func (p *Poisson) U(i, j int) float64 {
	return p.u[i][j]
}

func (p *Poisson) Rho(i, j int) float64 {
	x, y := float64(i)-N, float64(j)-N
	d := p.d
	x0 := p.x0
	return math.Exp(-(math.Pow((x-x0), 2)+y*y)/(d*d)) - math.Exp(-(math.Pow((x+x0), 2)+y*y)/(d*d))
}

func (p *Poisson) OptimizeAt(i, j int) {
	var s [4]sdelta
	for k, delta := range p.optimizationDeltas {
		s[k].delta = delta
		s[k].s = p.SLocal(i, j, delta)
	}

	s[3].delta = .25 * (3*s[0].s - 4*s[1].s + s[2].s) / (s[0].s - 2*s[1].s + s[2].s)
	s[3].s = p.SLocal(i, j, s[3].delta)

	sort.Slice(s[:], func(i, j int) bool {
		return s[i].s < s[j].s
	})

	p.u[i][j] += s[0].delta
}

// Optimize does one iteration (optimizeAt for each point). It returns total S value.
func (p *Poisson) Optimize() float64 {
	for i := 1; i <= 2*N+1-2; i++ {
		for j := 1; j <= 2*N+1-2; j++ {
			p.OptimizeAt(i, j)
		}
	}

	return p.S()
}

func (p *Poisson) SLocal(i, j int, delta float64) float64 {
	uij := p.U(i, j) + delta
	a1 := p.U(i-1, j) + p.U(i+1, j) + p.U(i, j-1) + p.U(i, j+1)
	a1 *= 2 * uij
	a2 := 4 * uij * uij
	a3 := p.Rho(i, j) * uij * math.Pow(p.dx, 2)

	return -0.5*(a1-a2) - a3
}

func (p *Poisson) S() float64 {
	var result float64
	dx2 := math.Pow(p.dx, 2)
	for i := 1; i <= 2*N+1-2; i++ {
		for j := 1; j <= 2*N+1-2; j++ {
			a1 := .5 * p.U(i, j) * (p.U(i+1, j) + p.U(i-1, j) - 2*p.U(i, j))
			a1 /= dx2
			a2 := .5 * p.U(i, j) * (p.U(i, j+1) + p.U(i, j-1) - 2*p.U(i, j))
			a2 /= dx2
			a3 := p.Rho(i, j) * p.U(i, j)
			result += (a1 + a2 + a3) * dx2
		}
	}
	return -result
}
