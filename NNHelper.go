package ml

func NNHelperConvertMultiArrayToSingleArray(multiArr [][]float64) []float64 {

	width := len(multiArr)
	height := len(multiArr[0])

	singleArr := make([]float64, width * height)

	i := 0
	for x:=0; x < width; x++ {
		for y:=0; y < height; y++ {
			singleArr[i] = multiArr[x][y]
			i++
		}
	}

	return singleArr
}

func NNHelperGetExpectedOutputs(index, total int) []float64 {
	outputs := make([]float64, total)
	for i:=0; i < total; i++ {
		outputs[i] = 0
	}

	outputs[index] = 1

	return outputs
}

func NNHelperGetBestIndexAndVaue(outputs []float64) (int, float64) {
	length := len(outputs)

	bestIndex := 0

	for i:=1; i < length; i++ {
		if outputs[i] > outputs[bestIndex] {
			bestIndex = i
		}
	}

	return bestIndex, outputs[bestIndex]
}