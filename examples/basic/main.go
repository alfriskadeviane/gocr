package main

import (
	"os"

	"github.com/alfriskadeviane/gocr"
)

var (
	samplePath = func() string {
		d, _ := os.Getwd()
		return d
	}() + "/../../train_data/"
)

var (
	modelPath = func() string {
		d, _ := os.Getwd()
		return d
	}() + "/../../model/"
)

func main() {
	// Downloading english font dataset and index.csv
	d, _ := os.Getwd()

	// Load the sample data and save it in .cbor file
	err := gocr.Train(samplePath+"sample2/", modelPath+"sample2/")
	if err != nil {
		panic(err)
	}

	// Read the image
	image, err := gocr.ReadImage(d + "/imagetext_2.png")
	if err != nil {
		panic(err)
	}

	im := gocr.ImageToGraysclaeArray(image)
	gocr.ImageMatrixToImage(im, d+"/result/image.png", 255)

}
