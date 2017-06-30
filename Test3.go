package main

import (
	"fmt"
	"math/rand"
	"time"
)

var net = NewNeuralNetwork(729, 10)

func GetValues(fileName string) ([][]float64, [][]float64) {
	folders := ReadingDirectorySearchFiles(fileName)

	totalOutputCount := len(folders)

	var inputs, outputs [][]float64

	for i, folder := range folders {
		imagesNames := ReadingDirectorySearchFiles(fileName + "/" + folder)

		for _, imageName := range imagesNames {

			image := ReadingImageGetImage(fileName + "/" + folder + "/" + imageName)
			newInput := NNHelperConvertMultiArrayToSingleArray(ReadingImageGetChangeValues(image))

			if len(newInput) == 729 {//9801 {
				newOutput := NNHelperGetExpectedOutputs(i, totalOutputCount)

				inputs = append(inputs, newInput)
				outputs = append(outputs, newOutput)
			} else {
				fmt.Println("ERROR: New input size is", len(newInput), "and not 729")
			}
		}
	}

	return inputs, outputs
}

func train(times int) {
	inputs, outputs := GetValues("mnist_png/training")

	fmt.Println("Got", len(inputs), "values")

	/*
	for i:=0; i < times; i++ {
		net.TrainMultiple(inputs, outputs)
		fmt.Println("Trained", i, "out of", times)
	}
	*/

	length := len(inputs)

	s1 := rand.NewSource(time.Now().UnixNano())//(1)//time.Now().UnixNano())
	r1 := rand.New(s1)

	for i:=0; i < times; i++ {
		randomIndex := r1.Intn(length)
		net.TrainSingle(inputs[randomIndex], outputs[randomIndex])
		fmt.Println("Trained", i, "out of", times, " ", float64(i) / float64(times), "%")
	}
}

func test() {
	fmt.Println("Testing")
	inputs, outputs := GetValues("mnist_png/training") //testing")
	length := len(inputs)

	fmt.Println("Got Test Values")

	correct:=0

	for i:=0; i < length; i++ {
		botOutPut := net.Output(inputs[i])
		fmt.Println(i, "/", length, "\t", botOutPut, " == ", outputs[i])
		index1, _ := NNHelperGetBestIndexAndVaue(outputs[i])
		index2, _ := NNHelperGetBestIndexAndVaue(botOutPut)

		if index1 == index2 {
			correct++
		}
	}

	percentRight := float64(correct) / float64(length)
	fmt.Println("Bot got,", correct, "out of", length, "correct. Thats", percentRight, "%")
}

func main() {

	fmt.Println("Starting:")

	train(1000000)

	test()
}
