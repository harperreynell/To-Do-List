package listHandling

import (
	"fmt"
)

type Todos struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Todo   string `json:"todo"`
}

func newTodo(task string, id int) Todos {
	todo := Todos{
		Todo:   task,
		Id:     id,
		Status: "opened"}

	return todo
}

func ChangeStatus(task *Todos) {
	task.Status = "closed"
}

func appendTask(todoList []Todos, todo string) []Todos {
	arr := append(todoList, newTodo(todo, len(todoList)))

	return arr
}

func TaskByID(todoList []Todos, id int) *Todos {
	return &todoList[id]
}

func PrintTodos(todoList []Todos) {
	var status string
	for i := range todoList {
		if todoList[i].Status == "closed" {
			status = "✓"
		} else {
			status = "✗"
		}

		fmt.Printf("[%s] id : %d, task: %s\n", status, todoList[i].Id, todoList[i].Todo)
	}
}

func TaskById(todoList []Todos, id int) *Todos {
	return &todoList[id]
}

func ReadTask(todoList *[]Todos) {
	var task string
	fmt.Scanln(&task)
	*todoList = appendTask(*todoList, task)

	//return todoList
}
