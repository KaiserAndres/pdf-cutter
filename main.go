package main

import (
	"os"
	"io/ioutil"
)

/// os.Args[1] is the input pdf
/// os.Args[2] is the output pdf
func main() {

	uncutDir := ExtractImages(os.Args[1])

	imageNames, err := ioutil.ReadDir(uncutDir)
	if err != nil {
		panic(err)
	}

	os.Mkdir("cut", 0777)
	cutDir := "cut" + string(os.PathSeparator)
	pageNumber := 0
	for _, file := range imageNames {
		DividePicture(uncutDir, file, cutDir, pageNumber)
		pageNumber += 2
	}

	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}

	pdf := JoinIMages(cutDir)
	pdf.OutputAndClose(outputFile)

	defer func() {
		os.RemoveAll(uncutDir)
		outputFile.Close()
		os.RemoveAll(cutDir)
	}()
}
