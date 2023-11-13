package listHandling

import (
	"fmt"
	"strconv"
)

type Todos struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Todo   string `json:"todo"`
}

func NewTodo(task string, id int) Todos {
	todo := Todos{
		Todo:   task,
		Id:     id,
		Status: "opened"}

	return todo
}

func ChangeStatus(task *Todos, status string) {
	task.Status = status
}

func appendTask(todoList []Todos, todo string) []Todos {
	arr := append(todoList, NewTodo(todo, len(todoList)))

	return arr
}

func TaskByID(todoList []Todos, id int) *Todos {
	return &todoList[id]
}

func TaskId(todoList []Todos, id int) int {
	for i := 0; i < len(todoList); i++ {
		if todoList[i].Id == id {
			return i
		}
	}

	return -1
}

func DeleteTaskByID(todoList []Todos, id int) []Todos {
	copy(todoList[id:], todoList[id+1:])
	return todoList[:len(todoList)-1]
}

func PrintTodos(todoList []Todos) string {
	var status, text string
	text = ""
	for i := range todoList {
		if todoList[i].Status == "closed" {
			status = "✓"
		} else {
			status = "✗"
		}
		id := strconv.Itoa(todoList[i].Id)
		text += "[" + status + "] id : " + id + ", task: " + todoList[i].Todo + "\n"
		//t.Client()
		//log.Println(text)
		//fmt.Printf("[%s] id : %d, task: %s\n", status, todoList[i].Id, todoList[i].Todo)
	}

	return text
}

func ReadTask(todoList *[]Todos) {
	var task string
	_, _ = fmt.Scanln(&task)
	*todoList = appendTask(*todoList, task)

	//return todoList
}
