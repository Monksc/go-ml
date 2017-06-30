package main

const LEARNING_RATE = 0.2 // should be 0.2

type Layer struct {
	neurons []Neuron
}

func NewLayer(currentLayerSize, nextLayerSize int, function ActivationFunction) Layer {
	neurons := make([]Neuron, currentLayerSize + 1)

	for i:= 0; i < currentLayerSize; i++ {
		neurons[i] = NewNeuron(nextLayerSize, function)
	}
	neurons[currentLayerSize] = NewNeuron(nextLayerSize, function)

	return Layer{neurons}
}

func (currentLayer *Layer) FeedForward(nextLayer *Layer) {
	cLength := len(currentLayer.neurons)
	nLength := len(currentLayer.neurons[0].weights)
	for n:=0; n < nLength; n++ {
		newValue := 0.0
		for c:=0; c < cLength; c++ {
			newValue += currentLayer.neurons[c].Output(n)
		}

		nextLayer.neurons[n].SetValue(newValue)
	}
}

func (l *Layer) GetValues() []float64 {
	count := len(l.neurons) - 1
	values := make([]float64, count)
	for i:=0; i<count; i++ {
		values[i] = l.neurons[i].value
	}

	return values
}

func (l *Layer) SetValues(newValues []float64) {
	count := len(newValues)
	for i:=0; i < count; i++ {
		l.neurons[i].value = newValues[i]
	}
}

func (cLayer *Layer) Train(nLayer *Layer) {
	currentNeuronCount := len(cLayer.neurons)
	weightsCount := len(cLayer.neurons[0].weights)
	for c:=0; c < currentNeuronCount; c++ {

		newChange := 0.0
		for n:=0; n < weightsCount; n++ {
			newChangeInWeight := (nLayer.neurons[n].totalChangeWeight / float64(currentNeuronCount)) * cLayer.neurons[c].value * LEARNING_RATE
			cLayer.neurons[c].weights[n] += newChangeInWeight
			newChange += newChangeInWeight
		}

		cLayer.neurons[c].totalChangeWeight = newChange
	}
}

func (l *Layer) GetWeights() [][]float64 {
	neuronsCount := len(l.neurons)
	weights := make([][]float64, neuronsCount)

	for i:=0; i < neuronsCount; i++ {
		weights[i] = l.neurons[i].weights
	}

	return weights
}

/**
 * Should only be called in last array
 * @param chanegWeights []float64 should be error or how far off the Network results are from expected results
 */
func (l *Layer) SetTotalChangeWeights(changeWeights []float64) {
	length := len(changeWeights)
	for i:=0; i < length; i++ {
		l.neurons[i].totalChangeWeight = changeWeights[i]
	}
}