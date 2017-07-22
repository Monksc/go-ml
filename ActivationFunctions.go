package ml

import (
	"math"
	"math/rand"
	"strconv"
)

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

func ActivationFunctionGetRawValueOfActivationFunction(function ActivationFunction) int {
	functions := []ActivationFunction{ActivationFunctionRelu, ActivationFunctionLeakyRelu, ActivationFunctionLogistic}

	indexes := []int{0,1,2}

	//didChange := true

	for len(functions) > 1 { //&& didChange {

		//didChange = false
		randomValue := rand.Float64()

		for i := 0; i < len(functions); i++ {
			if functions[i](randomValue) != functions[i](randomValue) {

				functionReplacement := functions[:i]
				indexesReplacement := indexes[:i]
				for j := i+1; j < len(functions); j++ {
					functionReplacement = append(functionReplacement, functions[j])
					indexesReplacement = append(indexesReplacement, indexes[j])
				}

				functions = functionReplacement
				indexes = indexesReplacement

				//didChange = true
			}
		}
	}

	if len(indexes) > 0 {
		return indexes[0]
	} else {
		return -1
	}
}

func ActivationFunctionGetActivationFunctionWithRawValue(rawValue int) ActivationFunction {
	switch rawValue {
	case 0:
		return ActivationFunctionRelu
	case 1:
		return ActivationFunctionLeakyRelu
	case 2:
		return ActivationFunctionLogistic
	default:
		message := "github.monksc.ml ActivationFunction.go ActivationFunctionGetActivationFunctionWithRawValue cant get ActivationFunction with rawValue of %v" + strconv.Itoa(rawValue)
		panic(message)
	}
}
