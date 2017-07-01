package ml

type ConvolutionalNeuralNetwork struct {
	width, height int
	neuralNetwork NeuralNetwork
}

func NewConvolutionalNeuralNetwork(outputSize, width, height int) ConvolutionalNeuralNetwork {
	neuralNetwork := NewNeuralNetwork(width * height, outputSize)
	return ConvolutionalNeuralNetwork{width, height, neuralNetwork}
}

/**
 * Gets output of a section of the total frame of the image
 */
func (c *ConvolutionalNeuralNetwork) GetSingleFrameOutput(image [][]float64) []float64 {
	inputs := NNHelperConvertMultiArrayToSingleArray(image)

	return c.neuralNetwork.Output(inputs)
}

/**
 * Gets max output of the full screen
 * You will get one place with the most likely to be object
 */
func (c *ConvolutionalNeuralNetwork) GetOutputOfScreen(screen [][]float64) (int, float64, int, int) {

	bestIndexValue := 0
	bestPercent := 0.0
	bestX := 0
	bestY := 0

	screenWidth := len(screen)
	screenHeight := len(screen[0])
	for x:=0; x <= screenWidth - c.width; x++ {
		for y:=0; y <= screenHeight - c.height; y++ {

			inputs := convertSectionOfAMultiArrayToSingleArray(screen, x, y, c.width, c.height)
			outputs := c.neuralNetwork.Output(inputs)
			index, value := getBestIndexAndValueOfArray(outputs)

			if value > bestPercent {
				bestIndexValue = index
				bestPercent = value
				bestX = x
				bestY = y

				if value == 1.0 {
					return bestIndexValue, bestPercent, bestX, bestY
				}
			} else if value < 0.1 {
				//fmt.Println("BLOOOOOOOOOOOOOBER ", value)
				y +=5
			}
		}
	}

	return bestIndexValue, bestPercent, bestX, bestY
}

/**
 * Trains On a section of the total frame
 */
func (c *ConvolutionalNeuralNetwork) Train(inputs [][]float64, expectedOutputs []float64) {
	singleInputs := NNHelperConvertMultiArrayToSingleArray(inputs)

	c.neuralNetwork.TrainSingle(singleInputs, expectedOutputs)
}

/* MARK: Private Functions */

func convertSectionOfAMultiArrayToSingleArray(multiArr [][]float64, x, y, width, height int) []float64 {
	returnFrame := make([]float64, width * height)

	i:=0
	for additonX:=0; additonX < width; additonX ++ {
		for additonY:=0; additonY < height; additonY++ {

			returnFrame[i] = multiArr[x][y]
			i++
		}
	}

	return returnFrame
}

func getBestIndexAndValueOfArray(arr []float64) (int, float64) {
	length := len(arr)

	bestIndex := 0
	bestValue := 0.0

	for i:=0; i < length; i++ {
		if arr[i] > bestValue {
			bestValue = arr[i]
			bestIndex = i
		}
	}

	return bestIndex, bestValue
}
