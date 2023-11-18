package web

import (
	"html/template"
	"net/http"
	"os"
	l "todolist/listHandling"
)

var tpl = template.Must(template.ParseFiles("html/index.html"))

type Todo struct {
	Item string
	Done bool
}

func newTodo(task string, status string) Todo {
	if status == "closed" {
		return Todo{
			Item: task,
			Done: true}
	}

	return Todo{
		Item: task,
		Done: false,
	}
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	var Todos []Todo
	todoList := l.ReadJsonFromFile(1)
	for i := 0; i < len(todoList); i++ {
		Todos = append(Todos, newTodo(todoList[i].Todo, todoList[i].Done))
	}
	data := PageData{
		Title: "Your list",
		Todos: Todos,
	}

	tpl.Execute(w, data)
}

func Create() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/todo", todo)

	http.ListenAndServe(":"+port, mux)
}
