package listHandling

import (
	"strconv"
)

type Todos struct {
	Id   int    `json:"id"`
	Done string `json:"status"`
	Todo string `json:"todo"`
}

func NewTodo(task string, id int) Todos {
	todo := Todos{
		Todo: task,
		Id:   id,
		Done: "opened"}

	return todo
}

func ChangeStatus(task *Todos, status string) {
	task.Done = status
}

func appendTask(todoList []Todos, todo string) []Todos {
	arr := append(todoList, NewTodo(todo, len(todoList)))

	return arr
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
		if todoList[i].Done == "closed" {
			status = "✓"
		} else {
			status = "✗"
		}
		id := strconv.Itoa(todoList[i].Id + 1)
		text += "[" + status + "] id : " + id + ", task: " + todoList[i].Todo + "\n"
	}

	return text
}

func ReadTask(todoList *[]Todos, task string) {
	*todoList = appendTask(*todoList, task)

	//return todoList
}
