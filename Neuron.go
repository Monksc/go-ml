package ml

// "time"
import (
	"math/rand"
	"time"
)

type Neuron struct {
	value float64
	weights [] float64
	activationFunction ActivationFunction
	totalChangeWeight float64
}

func NewNeuron(nextLayerSize int, activationFunction ActivationFunction) Neuron {
	weights := make([]float64, nextLayerSize)

	s1 := rand.NewSource(time.Now().UnixNano()) //(1)//time.Now().UnixNano())
	r1 := rand.New(s1)

	for i:=0; i < nextLayerSize; i++ {
		weights[i] = r1.Float64()
	}

	return Neuron{1,weights,activationFunction, 0}
}

func (n *Neuron) Output(index int) float64 {
	return n.weights[index] * n.value
}

func (n* Neuron) SetValue(newValue float64) {
	n.value = n.activationFunction(newValue)
}
