package main

import "fmt"

func train(network *ConvolutionalNeuralNetwork) {
	folders := ReadingDirectorySearchFiles("TrainingDataSet")

	totalOutputCount := len(folders)

	fmt.Println("Training With Dataset Of ", folders)

	for i, folder := range folders {
		imagesNames := ReadingDirectorySearchFiles("TrainingDataSet/" + folder)

		fmt.Println("Training:", folder)

		for _, imageName := range imagesNames {

			image := ReadingImageGetImage("TrainingDataSet/" + folder + "/" + imageName)
			inputs := ReadingImageGetChangeValues(image)

			if len(inputs) == 99 && len(inputs[0]) == 99 {

				outputs := NNHelperGetExpectedOutputs(i, totalOutputCount)

				network.Train(inputs, outputs)
			}
		}
	}
}

func test(network ConvolutionalNeuralNetwork) {
	testFolders := ReadingDirectorySearchFiles("TestDataSet")

	for _, folder := range testFolders {
		imagesNames := ReadingDirectorySearchFiles("TestDataSet/" + folder)

		fmt.Println("Testing:", folder)

		for _, imageName := range imagesNames {

			image := ReadingImageGetImage("TestDataSet/" + folder + "/" + imageName)
			inputs := ReadingImageGetChangeValues(image)


			if len(inputs) == 99 && len(inputs[0]) == 99 {
				fmt.Println("HERE POTATOE vcsd")
				outputs := network.GetSingleFrameOutput(inputs)
				guessedIndex, guessedChance := NNHelperGetBestIndexAndVaue(outputs)

				fmt.Println("Outputs:", outputs)
				fmt.Println("Was ", folder, ". Guessed", testFolders[guessedIndex], "with", guessedChance, "% confidence")
			} else {

				//index, percent, _, _ := network.GetOutputOfScreen(inputs)
				//fmt.Println(folder, " == ", testFolders[index], "with", percent)
			}
		}
	}
}

func main() {
	cnn := NewConvolutionalNeuralNetwork(2, 100, 100)

	for i:=0; i < 10; i++ {
		fmt.Println("Training", i, "out of", 25)
		train(&cnn)
		fmt.Println("\n")
	}

	fmt.Println("\n\n\nTesting Data:")

	test(cnn)
}
