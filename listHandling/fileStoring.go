package listHandling

import (
	"encoding/json"
)

func toJson(data Todos) []byte {
	b, err := json.Marshal(data)
	if err != nil {
	}

	return b
}

func fromJson(data []byte) Todos {
	var task Todos
	err := json.Unmarshal(data, &task)
	if err != nil {
	}

	return task
}

func Store() {
	data := Todos{
		Id:     1,
		Status: "opened",
		Todo:   "test1",
	}

	b := toJson(data)
	task := fromJson(b)

	todoList := make([]Todos, 0)
	todoList = append(todoList, task)
	PrintTodos(todoList)
}
