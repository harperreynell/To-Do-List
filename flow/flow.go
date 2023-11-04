package flow

import (
	"fmt"
	l "todolist/listHandling"
)

func printMenu() {
	fmt.Println("\t1 Print tasks")
	fmt.Println("\t2 Add task")
	fmt.Println("\t3 Change task status")
	fmt.Println("\t99 Exit")
}

func taskByID(todoList []l.Todos, id int) *l.Todos {
	return &todoList[id]
}

func Flow() {
	todos := make([]l.Todos, 0)
	var choice, id int = 1, 0

	printMenu()

	for choice != 99 {
		fmt.Print("Option #> ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:
			l.PrintTodos(todos)
		case 2:
			fmt.Println("Type in what you have to do:")
			l.ReadTask(&todos)
		case 3:
			fmt.Println("Type in id of task(id can be found by executing 1 option):")
			_, err := fmt.Scanln(&id)
			if err != nil || id >= len(todos) {
				fmt.Println("No such id")
				continue
			}
			l.ChangeStatus(taskByID(todos, id))
		}
	}
}
