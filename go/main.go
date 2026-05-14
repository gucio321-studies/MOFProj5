package main

import (
	"fmt"
	"image"

	"github.com/AllenDang/giu"
	"github.com/disintegration/imaging"
	"github.com/gucio321-studies/MOFProj5/go/pkg"
	"github.com/kpango/glg"
)

const (
	dx    = 1
	x0    = 4
	d     = 4
	dFall = .001
)

var betaSlider int32

func loop(ss []float64, uImg image.Image, betas []float64, ssFall [][]float64, fallImages []image.Image, ssRand []float64, imgRand []image.Image) {
	giu.SingleWindow().Layout(
		giu.TabBar().TabItems(
			giu.TabItem("Zadanie 1").Layout(
				giu.Plot("Zależność wartości S od numeru iteracji").Plots(
					giu.Line("S(it)", ss),
				).XAxeFlags(giu.PlotAxisFlagsAutoFit).YAxeFlags(giu.PlotAxisFlagsAutoFit, 0, 0),
				giu.Align(giu.AlignCenter).To(
					giu.ImageWithRgba(uImg),
				),
			),
			giu.TabItem("Zadanie 2").Layout(
				giu.Plot("Zależność wartości S od numeru iteracji dla wybranych współczynników Beta").Plots(
					giu.Custom(func() {
						for i, beta := range betas {
							giu.Line(fmt.Sprintf("Beta=%.2f", beta), ssFall[i]).Plot()
						}
					}),
					giu.Line("S dla minimalizacji bezpośredniej", ss),
				).XAxeFlags(giu.PlotAxisFlagsAutoFit).YAxeFlags(giu.PlotAxisFlagsAutoFit, 0, 0),
				giu.Align(giu.AlignCenter).To(
					giu.SliderInt(&betaSlider, 0, int32(len(betas)-1)).Format(fmt.Sprintf("Beta: %.2f", betas[betaSlider])),
					giu.ImageWithRgba(fallImages[betaSlider]),
				),
			),
			giu.TabItem("Zadanie 5").Layout(
				giu.Plot("Zależność wartości S od numeru iteracji dla optymalizacji MonteCarlo").Plots(
					giu.Line("S dla minimalizacji losowej", ssRand),
					giu.Line("S dla minimalizacji bezpośredniej", ss),
				).XAxeFlags(giu.PlotAxisFlagsAutoFit).YAxeFlags(giu.PlotAxisFlagsAutoFit, 0, 0),
				giu.Custom(func() {
					for _, img := range imgRand {
						giu.Align(giu.AlignCenter).To(
							giu.ImageWithRgba(img),
							giu.Labelf("Obraz u dla iteracji %d", len(imgRand)),
						).Build()
					}
				}),
			),
		),
	)
}

func main() {
	glg.Info("MOF5: Równanie Poissona: minimalizacja, optymalizacja gradientowa i MonteCarlo")

	glg.Info("Running task 1")
	p := pkg.NewPoisson(d, x0, dx, [3]float64{0, 0.5, 1}, dFall, 0)
	const nIter = 500
	ss := make([]float64, nIter)
	for i := range nIter {
		ss[i] = p.Optimize(p.OptimizeAt)
	}

	uImg := p.GetUMap()
	uImg = imaging.Resize(uImg, 512, 512, imaging.NearestNeighbor)
	glg.Success("Task 1 complete!")

	glg.Info("Running task 2")
	betas := []float64{0.1, 0.3, 0.4, 0.49}
	ssFalls := make([][]float64, 0)
	fallImages := make([]image.Image, 0)
	for _, beta := range betas {
		glg.Debugf("Running for beta=%.2f", beta)
		p = pkg.NewPoisson(d, x0, dx, [3]float64{0, 0.5, 1}, dFall, beta)
		ssFall := make([]float64, nIter)
		for i := range nIter {
			ssFall[i] = p.Optimize(p.OptimizeFallAt)
		}

		ssFalls = append(ssFalls, ssFall)
		uImg := p.GetUMap()
		uImg = imaging.Resize(uImg, 512, 512, imaging.NearestNeighbor)
		fallImages = append(fallImages, uImg)
	}

	glg.Success("Task 2 complete!")

	glg.Info("Running task 3")
	const r = 0.1

	p = pkg.NewPoisson(d, x0, dx, [3]float64{0, 0.5, 1}, dFall, 0)
	imgRand := make([]image.Image, 0)
	ssRand := make([]float64, nIter)
	for i := range nIter {
		ssRand[i] = p.Optimize(func(i, j int) {
			p.OptimizeRandAt(i, j, r)
		})
		if i%100 == 0 {
			glg.Debugf("Map for iteration %d saved", i)
			uImg := p.GetUMap()
			uImg = imaging.Resize(uImg, 512, 512, imaging.NearestNeighbor)
			imgRand = append(imgRand, uImg)
		}
	}

	glg.Success("Task 3 complete!")

	glg.Info("Running visualization with giu.")
	wnd := giu.NewMasterWindow("Poisson", 640, 480, 0)
	wnd.SetStyle(giu.LightTheme())
	wnd.Run(func() {
		loop(ss, uImg, betas, ssFalls, fallImages, ssRand, imgRand)
	})
}
