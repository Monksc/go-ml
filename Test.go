package main

import (
	"fmt"
)

func main() {

	inputs := [][]float64 {
		[]float64{0.0,0.0},
		[]float64{1.0,0.0},
		[]float64{0.0,1.0},
		[]float64{1.0,1.0},
	}

	outputs := [][]float64 {
		[]float64{1.0},
		[]float64{0.0},
		[]float64{0.0},
		[]float64{1.0},
	}

	net := NewNeuralNetwork([]int{2,2,1} ...)

	for i:=0; i < 10000; i++ {
		net.TrainMultiple(inputs, outputs)
	}

	for i:=0; i < len(inputs); i++ {
		fmt.Printf("%v \t %v\n", net.Output(inputs[i]), outputs[i])
	}

	weights := net.GetWeights()

	for l:=0; l < len(weights); l++ {
		for n:=0; n < len(weights[l]); n++ {
			for w:=0; w < len(weights[l][n]); w++ {
				fmt.Printf("%.2f ", weights[l][n][w])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
