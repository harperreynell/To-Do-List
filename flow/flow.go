package flow

// For console application

import (
	"fmt"
	l "todolist/listHandling"
)

func printMenu() {
	fmt.Println("\t1 Print tasks")
	fmt.Println("\t2 Add task")
	fmt.Println("\t3 Delete task")
	fmt.Println("\t4 Change task status")
	fmt.Println("\t5 Print help")
	fmt.Println("\t99 Exit")
}

func Flow() {
	todos := make([]l.Todos, 0)
	var choice, id int = 1, 0
	var status string

	printMenu()

	for choice != 99 {
		fmt.Print("Option #> ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:
			todos = l.ReadJsonFromFile()
			l.PrintTodos(todos)
		case 2:
			fmt.Println("Type in what you have to do:")
			l.ReadTask(&todos)
			l.WriteJsonToFile(todos)
			fmt.Println("Task was added successfully")
		case 3:
			todos = l.ReadJsonFromFile()
			_, err := fmt.Scanln(&id)
			if err != nil || id >= len(todos) {
				fmt.Println("No such id")
				continue
			}
			todos = l.DeleteTaskByID(todos, id)
			l.WriteJsonToFile(todos)
			fmt.Println("Task was deleted successfully")
		case 4:
			fmt.Println("Type in id of task(id can be found by executing 1 option):")
			_, err := fmt.Scanln(&id)
			if err != nil || id >= len(todos) {
				fmt.Println("No such id")
				continue
			}

			fmt.Println("Type in status [opened/closed]:")
			_, _ = fmt.Scanln(&status)
			if (status != "opened") && (status != "closed") {
				fmt.Println("Unable to set status")
			}
			l.ChangeStatus(l.TaskByID(todos, id), status)
			l.WriteJsonToFile(todos)
			fmt.Println("Status changed successfully")
		case 5:
			printMenu()
		}
	}
}
