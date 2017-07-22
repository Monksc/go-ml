package ml

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DataBaseUpdateNeuralNetwork(network DataBaseNeuralNetwork) int64 {
	db, err := sql.Open("mysql", "root:Password@tcp(localhost:3306)/neuralNetwork")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil{
		panic(err)
	}

	stmt, err := db.Prepare("update Network set layersCount = ? where id = ?")
	checkErr(err)

	res, err := stmt.Exec(len(network.layers), network.id)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	return id
}

func GetMessages() []*Message {
	//db, err := sql.Open("mysql", "root:Password@unix(192.168.0.20:3306)/cam")
	db, err := sql.Open("mysql", "root:Password@tcp(localhost:3306)/cam")

	if err != nil{

		fmt.Println(err)
		return []*Message {}
	}
	defer db.Close()

	//check to see if we have a connection
	err = db.Ping()
	if err != nil{

		fmt.Println("There was an error", err)
	} else{

		rows, err := db.Query("SELECT * FROM chatRoom")
		if err != nil {
			panic(err)
		}

		messages := []*Message {}
		//and finally add this object
		for rows.Next() {
			var id int
			var message string

			err = rows.Scan(&id, &message)
			if err != nil {
				panic(err)
			}
			fmt.Println(id)
			fmt.Println(message)

			messages = append(messages, &Message{id, message})
		}

		return messages
	}
	return []*Message {}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}