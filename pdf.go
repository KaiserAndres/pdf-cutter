package main

import (
	"log"
	"os"
	"github.com/MrSaints/go-ghostscript/ghostscript"
	"io/ioutil"
	"github.com/jung-kurt/gofpdf"
)

func ExtractImages(inputFile string) string {
	rev, err := ghostscript.GetRevision()
	if err != nil {
		panic(err)
	}
	log.Printf("Revision: %+v\n", rev)
	gs, err := ghostscript.NewInstance()

	if err != nil {
		panic(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	divider := string(os.PathSeparator)
	target := pwd + divider + "uncut"
	os.Mkdir(target, 0777)

	args := []string{
		"gs",
		"-sDEVICE=pngalpha",
		"-o",
		 target + divider + "page-%03d.png",
		"-r200",
		inputFile,
	}

	if err := gs.Init(args); err != nil {
		panic(err)
	}
	defer func() {
		gs.Exit()
		gs.Destroy()
	}()

	return target
}

func JoinIMages(path string) *gofpdf.Fpdf{

	imageNames , err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	pdf := gofpdf.New("P", "mm", gofpdf.PageSizeA4, "")

	for _, file := range imageNames {
		imageName := path + string(os.PathSeparator) + file.Name()
		if err != nil {
			panic(err)
		}

		fileHandler, err := os.Open(imageName)
		if err != nil {
			panic(err)
		}

		//f, _ := os.Open(imageName)
		//image, _ := png.Decode(f)
		width, height := pdf.GetPageSize()
		pdf.AddPage()
		pdf.ImageOptions(
			imageName,
			0, 0,
			width, height,
			false, gofpdf.ImageOptions{"png", true},
			0, "")
		pdf.GetImageInfo(imageName).SetDpi(300)
		fileHandler.Close()
	}

	return pdf
}