package flow

import (
	"fmt"
	l "todolist/listHandling"
)

func Flow() {
	todos := make([]l.Todos, 0)
	fmt.Println(todos)

	l.ReadTask(&todos)
	l.ChangeStatus(&todos[2])
	l.PrintTodos(todos)
}
