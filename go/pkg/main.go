package pkg

import (
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

func (p *Poisson) U(i, j int) float64 {
	return p.u[i][j]
}

func (p *Poisson) Rho(i, j int) float64 {
	x, y := float64(i), float64(j)
	d := p.d
	x0 := p.x0
	return math.Exp(-(math.Pow((x-x0), 2)+y*y)/(d*d)) - math.Exp(-(math.Pow((x+x0), 2)+y*y)/(d*d))
}

func (p *Poisson) OptimizeAt(i, j int) {
	var s [4]sdelta
	for k, delta := range p.optimizationDeltas {
		s[k].delta = delta
		s[k].s = p.S(i, j, delta)
	}

	s[3].delta = .25 * (3*s[0].s - 4*s[1].s + s[2].s) / (s[0].s - 2*s[1].s + s[2].s)
	s[3].s = p.S(i, j, s[3].delta)

	sort.Slice(s[:], func(i, j int) bool {
		return s[i].s < s[j].s
	})

	p.u[i][j] += s[0].delta
}

func (p *Poisson) Optimize() {
	for i := 1; i <= 2*N+1-2; i++ {
		for j := 1; j <= 2*N+1-2; j++ {
			p.OptimizeAt(i, j)
		}
	}
}

func (p *Poisson) S(i, j int, delta float64) float64 {
	uij := p.U(i, j) + delta
	a1 := p.U(i-1, j) + p.U(i+1, j) + p.U(i, j-1) + p.U(i, j+1)
	a1 *= 2 * uij
	a2 := 4 * uij * uij
	a3 := p.Rho(i, j) * uij * math.Pow(p.dx, 2)

	return -0.5*(a1-a2) - a3
}
