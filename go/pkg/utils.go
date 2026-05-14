package pkg

import (
	"image"
	"image/color"
)

func Float64ToGrayImage(data [2*N + 1][2*N + 1]float64) *image.Gray {
	height := len(data)
	width := len(data[0])

	// Find min/max
	min := data[0][0]
	max := data[0][0]

	for y := range data {
		for x := range data[y] {
			v := data[y][x]
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
		}
	}

	scale := 255.0
	if max != min {
		scale = 255.0 / (max - min)
	}

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := range data {
		for x := range data[y] {
			v := uint8((data[y][x] - min) * scale)
			img.SetGray(x, y, color.Gray{Y: v})
		}
	}

	return img
}
