package main

import "math"

type ActivationFunction func(float64) float64

func ActivationFunctionRelu(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}

func ActivationFunctionLeakyRelu(x float64) float64 {
	if x < 0 {
		return x / 5
	}
	return x
}

func ActivationFunctionLogistic(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}
