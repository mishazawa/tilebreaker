package main

import (
	"log"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"os"
)

func main () {
	err := os.Chdir("./testdata/input")
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	combineImage := image.NewNRGBA(image.Rect(0, 0, 1, 1))

	for _, file := range files {
		reader, err := os.Open(file.Name())
		if err != nil {
			log.Fatal(err)
		}

		currentImage, _, err := image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}

		boundsTo   := combineImage.Bounds()
		boundsFrom := currentImage.Bounds()

		aspect := boundsTo.Dy() / boundsTo.Dx()

		maxDx := Max(boundsFrom.Dx(), boundsTo.Dx())
		maxDy := Max(boundsFrom.Dy(), boundsTo.Dy())

		newWidth  := boundsTo.Dx()
		newHeight := boundsTo.Dy()

		var newRect image.Rectangle

		if aspect >= 3 {
			// put to right
			if maxDy > newHeight {
				newHeight = maxDy
			}
			newRect = image.Rect(0, 0, boundsTo.Dx() + boundsFrom.Dx(), newHeight)
		} else {
			// put to down
			if maxDx > newWidth {
				newWidth = maxDx
			}
			newRect = image.Rect(0, 0, newWidth, boundsTo.Dy() + boundsFrom.Dy())
		}

		combineImage = image.NewNRGBA(newRect)
		log.Println(file.Name(), combineImage.Bounds())
		reader.Close()
	}
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}
