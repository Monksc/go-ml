package ml

type NeuralNetwork struct {
	layers []Layer
}

type DataBaseNeuralNetwork struct {
	NeuralNetwork
	id int64
}

func NewNeuralNetwork(formation ...int) NeuralNetwork {
	formationCount := len(formation)
	layers := make([]Layer, formationCount)

	for i:=0; i < (formationCount - 1); i++ {
		layers[i] = NewLayer(formation[i], formation[i+1], ActivationFunctionRelu)
	}

	layers[formationCount - 1] = NewLayer(formation[formationCount - 1], 0, ActivationFunctionLogistic)

	return NeuralNetwork{layers}
}

func (n *NeuralNetwork) Output(input []float64) []float64 {
	n.layers[0].SetValues(input)
	length := len(n.layers) - 1
	for i:=0; i < length; i++ {
		n.layers[i].FeedForward(&n.layers[i+1])
	}

	return n.layers[length].GetValues()
}

func (n *NeuralNetwork) TrainSingle(inputs, expectedOutputs []float64) {
	realOutputs := n.Output(inputs)

	dif := differentsInArray(expectedOutputs, realOutputs)

	layersLength := len(n.layers)
	n.layers[layersLength - 1].SetTotalChangeWeights(dif)

	for i:=layersLength - 2; i >= 0; i-- {
		n.layers[i].Train(&n.layers[i+1])
	}
}

func (n *NeuralNetwork) TrainMultiple(inputs, expectedOutputs [][]float64) {
	length := len(inputs)
	for i:=0; i < length; i++ {
		n.TrainSingle(inputs[i], expectedOutputs[i])
	}
}

func (n *NeuralNetwork) GetWeights() [][][]float64 {
	layersCount := len(n.layers)
	weights := make([][][]float64, layersCount)

	for i:=0; i < layersCount; i++ {
		weights[i] = n.layers[i].GetWeights()
	}

	return weights
}

func differentsInArray(arr1, arr2 []float64) []float64 {

	length := len(arr1)
	dif := make([]float64, length)

	for i:=0; i < length; i++ {
		dif[i] = arr1[i] - arr2[i]
	}

	return dif
}
