package listHandling

import (
	"fmt"
)

type Todos struct {
	id     int
	status string
	todo   string
}

func newTodo(task string, id int) Todos {
	todo := Todos{
		todo:   task,
		id:     id,
		status: "opened"}

	return todo
}

func ChangeStatus(task *Todos) {
	task.status = "closed"
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
		if todoList[i].status == "closed" {
			status = "✓"
		} else {
			status = "✗"
		}

		fmt.Printf("[%s] id : %d, task: %s\n", status, todoList[i].id, todoList[i].todo)
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
