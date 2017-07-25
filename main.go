package ml

import (
	"fmt"
	"database/sql"
)


func getDatabase() *sql.DB  {
	db, err := sql.Open("mysql", "root:Password@tcp(localhost:3306)/neuralNetwork")

	if err != nil{
		fmt.Println(err)
		return nil
	}
	defer db.Close()

	//check to see if we have a connection
	err = db.Ping()
	if err != nil {

		fmt.Println("There was an error", err)
		return nil
	}

	return db
}

func train(network *DataBaseNeuralNetwork) {
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
		network.TrainMultiple(inputs, outputs)
	}
}

func output(network *DataBaseNeuralNetwork) {

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

	for i:=0; i < len(inputs); i++ {

		output := network.Output(inputs[i])
		fmt.Println(inputs[i], " -> ", output, "should be", outputs[i])
	}
}

func main() {

	db := getDatabase()

	nn := NewNeuralNetwork(2,2,1)
	net := DataBaseNeuralNetwork{nn, 1}

	insertNewNetwork := true
	if insertNewNetwork {
		net.InsertNeuralNetwork(db)
	}

	willTrain := true
	if willTrain {
		train(&net)
	}

	net.UpdateLayersWeights(db)

	output(&net)
}