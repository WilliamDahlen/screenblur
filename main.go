package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/kbinani/screenshot"
)

func capture() string {
	n := screenshot.NumActiveDisplays()
	fileName := "1"
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName = fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
	}
	return string(fileName)
}

func blur(file string) {
	src, err := imaging.Open(file)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	img1 := imaging.Blur(src, 10)
	dst := imaging.New(2560, 1440, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img1, image.Pt(0, 0))

	err = imaging.Save(dst, "lock.png")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func main() {
	blur(capture())
}
