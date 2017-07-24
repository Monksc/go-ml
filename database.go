package ml

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"fmt"
)

func getWeights(db *sql.DB, neuronId int64) []float64 {
	rows, err := db.Query("SELECT * FROM Weight where neuronId = " + strconv.Itoa(int(neuronId)))
	if err != nil {
		panic(err)
	}

	var weights []float64

	for rows.Next() {
		var id int
		var neuronId int
		var weightIndex int
		var value float64

		err = rows.Scan(&id, &neuronId, &weightIndex, &value)
		if err != nil {
			panic(err)
		}

		weights = append(weights, value)
	}

	return weights
}
func getNeurons(db *sql.DB, layerId int64) []Neuron {
	rows, err := db.Query("SELECT id, activationFunction FROM Neuron where layerId = " + strconv.Itoa(int(layerId)))
	if err != nil {
		panic(err)
	}

	var neurons [] Neuron

	for rows.Next() {
		var id int64
		var activationFunction int

		err = rows.Scan(&id, &layerId, &activationFunction)
		if err != nil {
			panic(err)
		}

		weights := getWeights(db, id)
		actualActivationFunction := ActivationFunctionGetActivationFunctionWithRawValue(activationFunction)

		neuron := Neuron{1,weights,actualActivationFunction, 0}
		neurons = append(neurons, neuron)
	}

	return neurons
}
func getLayers(db *sql.DB, networkId int64) []Layer {

	rows, err := db.Query("SELECT id FROM Layer where networkId = " + strconv.Itoa(int(networkId)))
	if err != nil {
		panic(err)
	}

	var layers [] Layer

	for rows.Next() {
		var id int64

		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}

		neurons := getNeurons(db, id)
		newLayer := Layer{neurons}

		layers = append(layers, newLayer)
	}

	return layers
}
func GetNetworkWithId(db *sql.DB, id int64) DataBaseNeuralNetwork {

	layers := getLayers(db, id)

	return DataBaseNeuralNetwork{NeuralNetwork{layers}, id}
}
func GetNetworks(db *sql.DB) []DataBaseNeuralNetwork{

	rows, err := db.Query("SELECT id FROM Network")
	if err != nil {
		panic(err)
	}

	var networks [] DataBaseNeuralNetwork

	for rows.Next() {
		var id int64

		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}


		newNetwork := GetNetworkWithId(db, id)
		networks = append(networks, newNetwork)
	}

	return networks
}

func (neuron *Neuron) updateWeights(db *sql.DB, neuronId int) {

	for index, weight := range neuron.weights {
		stmt, err := db.Prepare("update Weight set value = ? where neuronId = ? AND weightIndex = ?")
		checkErr(err)

		res, err := stmt.Exec(weight, neuronId, index)
		checkErr(err)

		_, err = res.LastInsertId()
		checkErr(err)
	}
}
func (layer *Layer) updateNeuronsWeights(db *sql.DB, layerId int) {

	for index, neuron := range layer.neurons {

		query := "select id from Neuron where layerId = " + strconv.Itoa(layerId) + " AND neuronIndex = " + strconv.Itoa(index
		rows, err := db.Query(query))

		if err != nil {
			fmt.Println(query)
			panic(err)
		}

		for rows.Next() {
			var id int

			err = rows.Scan(&id)
			if err != nil {
				panic(err)
			}

			neuron.updateWeights(db, id)
		}
	}

}
func (network *DataBaseNeuralNetwork) UpdateLayersWeights(db *sql.DB) {

	for index, layer := range network.layers {
		rows, err := db.Query("select id from Layer where networkId = " + strconv.Itoa(int(network.ID)) + " AND layerIndex = " + strconv.Itoa(index))

		if err != nil {
			panic(err)
		}


		for rows.Next() {
			var id int

			err = rows.Scan(&id)
			if err != nil {
				panic(err)
			}

			layer.updateNeuronsWeights(db, id)
		}
	}

}


func insertWeight(weightValue float64, db *sql.DB, neuronId int64, weightIndex int) {
	stmt, err := db.Prepare("insert Weight(neuronId, weightIndex, value) VALUES( ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(neuronId, weightIndex, weightValue)
	checkErr(err)

	_, err = res.LastInsertId()
	checkErr(err)
}
func (neuron *Neuron) insertNeuron(db *sql.DB, layerId int64, neuronIndex int) {
	stmt, err := db.Prepare("insert Neuron(layerId, neuronIndex, weightsCount, activationFunction) VALUES( ?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(layerId, neuronIndex, len(neuron.weights), ActivationFunctionGetRawValueOfActivationFunction(neuron.activationFunction))
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	for index, weight := range neuron.weights {
		insertWeight(weight, db, id, index)
	}
}
func (layer *Layer) insertLayer(db *sql.DB, networkId int64, layerIndex int) {
	stmt, err := db.Prepare("insert Layer(networkId, layerIndex, neuronsCount) VALUES( ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(networkId, layerIndex, len(layer.neurons))
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	for index, neuron := range layer.neurons {
		neuron.insertNeuron(db, id, index)
	}
}
func (network *DataBaseNeuralNetwork) InsertNeuralNetwork() {
	db, err := sql.Open("mysql", "root:Password@tcp(localhost:3306)/neuralNetwork")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}

	stmt, err := db.Prepare("insert Network VALUES(layersCount) ( ? )")
	checkErr(err)

	res, err := stmt.Exec(len(network.layers), network.ID)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	network.ID = id

	for index, layer := range network.layers {
		layer.insertLayer(db, id, index)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
