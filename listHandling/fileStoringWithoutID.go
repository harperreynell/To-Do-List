package listHandling

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJsonToFileWithoutID(data []Todos) {
	file, _ := os.Create("tasks.json")
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println(err)
		return
	}
}

func ReadJsonFromFileWithoutID() []Todos {
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
