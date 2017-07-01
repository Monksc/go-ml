package ml

import (
	"io/ioutil"
	"log"
)

func ReadingDirectorySearchFiles(dir string) []string {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string

	for _, file := range files {
		if file.Name()[0] != '.' {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames
}

/*
func main() {
	fmt.Println(ReadingDirectorySearchFiles("TrainingDataSet"))
}
*/
