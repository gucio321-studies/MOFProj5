package main

import (
	"fmt"

	"github.com/gucio321-studies/MMFProj5/go/pkg"
)

const (
	dx = 1
	x0 = 4
	d  = 4
)

func main() {
	p := pkg.NewPoisson(d, x0, dx, [3]float64{0, 0.5, 1})
	for range 500 {
		p.Optimize()
	}
	fmt.Println(p)
}
