package ch1

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func Sub4() {
	fmt.Print(" - Chapter 1.4\n\n")
	lissajousWithFilename("main/resources/ch14.gif")
}

func lissajousWithFilename(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lissajousWithFilename: %v\n", err)
		return
	}
	DefaultLissajous(f)
	f.Close()
}

type LissajousParameters struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

var DefaultLissajousParameters = LissajousParameters{5, 0.001, 100, 64, 8}

func DefaultLissajous(out io.Writer) {
	Lissajous(out, DefaultLissajousParameters)
}

func Lissajous(out io.Writer, parameters LissajousParameters) {
	var palette = []color.Color{color.White, color.Black}
	const (
		whiteIndex = 0 // first color in palette
		blackIndex = 1 // next color in palette
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: parameters.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < parameters.nframes; i++ {
		rect := image.Rect(0, 0, 2*parameters.size+1, 2*parameters.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(parameters.cycles)*2*math.Pi; t += parameters.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				parameters.size+int(x*float64(parameters.size)+0.5),
				parameters.size+int(y*float64(parameters.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, parameters.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
