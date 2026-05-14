package main

import (
	"github.com/AllenDang/giu"
	"github.com/gucio321-studies/MMFProj5/go/pkg"
)

const (
	dx = 1
	x0 = 4
	d  = 4
)

func loop(ss []float64) {
	giu.SingleWindow().Layout(
		giu.Plot("Ex1").Plots(
			giu.Line("SS", ss),
		),
	)
}

func main() {
	p := pkg.NewPoisson(d, x0, dx, [3]float64{0, 0.5, 1})
	const nIter = 500
	ss := make([]float64, nIter)
	for i := range nIter {
		ss[i] = p.Optimize()
	}

	wnd := giu.NewMasterWindow("Poisson", 640, 480, 0)
	wnd.Run(func() {
		loop(ss)
	})
}
