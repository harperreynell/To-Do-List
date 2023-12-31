package main

// For console application

import (
	"fmt"
	"log"
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

func main() {
	todos := make([]l.Todos, 0)
	var choice, id int = 1, 0
	var status, task string

	printMenu()

	for choice != 99 {
		fmt.Print("Option #> ")
		_, err := fmt.Scanln(&choice)
		if err != nil {
			return
		}

		switch choice {
		case 1:
			todos = l.ReadJsonFromFileWithoutID()
			l.PrintTodos(todos)
		case 2:
			fmt.Println("Type in what you have to do:")
			_, err := fmt.Scanln(&task)
			if err != nil {
				log.Fatal(err)
			}
			l.ReadTask(&todos, task)
			l.WriteJsonToFileWithoutID(todos)
			fmt.Println("Task was added successfully")
		case 3:
			todos = l.ReadJsonFromFileWithoutID()
			_, err := fmt.Scanln(&id)
			if err != nil || id >= len(todos) {
				fmt.Println("No such id")
				continue
			}
			todos = l.DeleteTaskByID(todos, id)
			l.WriteJsonToFileWithoutID(todos)
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
			l.ChangeStatus(&todos[id], status)
			l.WriteJsonToFileWithoutID(todos)
			fmt.Println("Status changed successfully")
		case 5:
			printMenu()
		}
	}
}
