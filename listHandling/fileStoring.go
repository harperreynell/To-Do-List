package listHandling

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func WriteJsonToFile(data []Todos, id int64) {
	file, _ := os.Create("tasks" + strconv.Itoa(int(id)) + ".json")
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println(err)
		return
	}
}

func ReadJsonFromFile(id int64) []Todos {
	file, err := os.Open("tasks" + strconv.Itoa(int(id)) + ".json")
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
