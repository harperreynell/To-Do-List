package listHandling

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJsonToFile(data []Todos) {
	file, _ := os.Create("tasks.json")
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println(err)
		return
	}
}

func ReadJsonFromFile() []Todos {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	decoder := json.NewDecoder(file)
	var taskList []Todos

	if err := decoder.Decode(&taskList); err != nil {
		fmt.Println(err)
		return nil
	}
	return taskList
}
