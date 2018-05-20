package main

import (
	"image/png"
	"os"
	"github.com/oliamb/cutter"
	"image"
	"fmt"
)

func CutImage(fileName string) []image.Image {
	fileHandler, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	doublePage, err := png.Decode(fileHandler)
	if err != nil {
		panic(err)
	}

	midWidth := doublePage.Bounds().Max.X / 2

	// Save these two into a directory. Number them properly.
	leftImage, err := cutter.Crop(doublePage, cutter.Config{
		Width:  midWidth,
		Height: doublePage.Bounds().Max.Y,
		Anchor: image.Point{0, 0},
	})

	if err != nil {
		panic(err)
	}

	rightImage, err := cutter.Crop(doublePage, cutter.Config{
		Width:  midWidth,
		Height: doublePage.Bounds().Max.Y,
		Anchor: image.Point{midWidth, 0},
	})

	if err != nil {
		panic(err)
	}

	defer func() {
		fileHandler.Close()
	}()

	return []image.Image{leftImage, rightImage}
}

func DividePicture(uncutDir string, file os.FileInfo, cutDir string, pageNumber int) {
	pair := CutImage(uncutDir + string(os.PathSeparator) + file.Name())
	fh1, err := os.Create(cutDir + "page " + fmt.Sprintf("%04d",pageNumber) + ".png")
	if err != nil {
		panic(err)
	}
	fh2, err := os.Create(cutDir + "page " + fmt.Sprintf("%04d", pageNumber+1) + ".png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(fh1, pair[0])
	if err != nil {
		panic(err)
	}
	err = png.Encode(fh2, pair[1])
	if err != nil {
		panic(err)
	}

	defer func() {
		fh1.Close()
		fh2.Close()
	}()
}