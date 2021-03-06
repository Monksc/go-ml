package ml

import (
	"image"
	"math"
	"os"
	"fmt"
	_ "image/jpeg"
	_ "image/gif"
	_ "image/png"
	_ "image/png"
    _ "image/color"
)

func getDifferanceInFourArrayValues(arr1, arr2, arr3, arr4 []float64) float64 {
	value := 0.0
	length := len(arr1)

	for i:=0; i < length; i++ {
		diff := math.Abs(arr1[i] - arr2[i]) + math.Abs(arr1[i] - arr3[i]) + math.Abs(arr3[i] - arr4[i]) + math.Abs(arr2[i] - arr4[i])
		value += diff * diff
	}

	return math.Sqrt(value)
}

func getDifferanceValues(arr1, arr2 []float64) float64 {
	value := 0.0
	length := len(arr1)

	for i:=0; i < length; i++ {
		diff := arr1[i] - arr2[i]
		value += diff * diff
	}

	return math.Sqrt(value)
}

func ReadingImageGetChangeValues(src image.Image) [][]float64 {
	width := src.Bounds().Size().X
	height := src.Bounds().Size().Y
	difScreen := make([][]float64, width - 1)

	for x:=0; x < width - 1; x++ {

		newCollum := make([]float64, height - 1)
		for y:=0; y < height - 1; y++ {

			red1, green1, blue1, alpha1 := src.At(x,y).RGBA()
			red2, green2, blue2, alpha2 := src.At(x+1,y).RGBA()
			red3, green3, blue3, alpha3 := src.At(x,y+1).RGBA()
			red4, green4, blue4, alpha4 := src.At(x+1,y+1).RGBA()

			color1 := []float64{float64(red1),float64(green1), float64(blue1), float64(alpha1)}
			color2 := []float64{float64(red2),float64(green2), float64(blue2), float64(alpha2)}
			color3 := []float64{float64(red3),float64(green3), float64(blue3), float64(alpha3)}
			color4 := []float64{float64(red4),float64(green4), float64(blue4), float64(alpha4)}

			newCollum[y] = getDifferanceInFourArrayValues(color1,color2,color3,color4)
		}

		difScreen[x] = newCollum
	}

	return difScreen
}

func ReadingImageGetImage(fileName string) image.Image {

	//fmt.Println("Finding Image:", fileName)

	infile, err := os.Open(fileName)
	if err != nil {
		// replace this with real error handling
		fmt.Println("ERROR HERE 1")
		panic(err.Error())
	}
	defer infile.Close()

	// Decode will figure out what type of image is in the file on its own.
	// We just have to be sure all the image packages we want are imported.
	src, _, err := image.Decode(infile)
	if err != nil {
		// replace this with real error handling
		fmt.Println("ERROR HERE 2", src)
		panic(err.Error())
	}

	return src
}
