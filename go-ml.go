<<<<<<< HEAD
package main

import "fmt"

func main() {
	nn := NewNeuralNetwork(2,2,1)

	inputs := [][]float64 {
		[]float64 {0,0},
		[]float64 {0,1},
		[]float64 {1,0},
		[]float64 {1,1},
	}

	outputs := [][]float64 {
		[]float64 {1},
		[]float64 {0},
		[]float64 {0},
		[]float64 {1},
	}


	for i:=0; i < 10000; i++ {
		nn.TrainMultiple(inputs, outputs)
	}

	for i:=0; i < len(inputs); i++ {

		output := nn.Output(inputs[i])
		fmt.Println(inputs[i], " -> ", output)
	}
}
=======
package ml
>>>>>>> bf8f895431f912a2d2ca74ecda426b678019461c
