package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Todo struct {
	Item string `json:"item"`
	Done bool   `json:"done"`
}

type PageData struct {
	Title string
	Todos []Todo
}

var (
	tpl      = template.Must(template.ParseFiles("html/index.html"))
	todoList []Todo
)

func newTodo(task string, status bool) Todo {
	return Todo{
		Item: task,
		Done: status,
	}
}

func addTask(w http.ResponseWriter, r *http.Request) {
	var task string
	todoList = ReadJsonFromFile()
	todoList = append(todoList, newTodo(task, false))

	//data := PageData{
	//	Title: "Your list",
	//	Todos: todoList,
	//}
	//
	//tpl.Execute(w, data)
}

func WriteJsonToFile(data []Todo) {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(data); err != nil {
		fmt.Println(err)
		return
	}
}

func ReadJsonFromFile() []Todo {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var todoList []Todo

	if err := decoder.Decode(&todoList); err != nil {
		fmt.Println(err)
		return nil
	}
	return todoList

}

func todo(w http.ResponseWriter, r *http.Request) {
	//WriteJsonToFile(Todos)
	Todos := ReadJsonFromFile()

	data := PageData{
		Title: "Your list",
		Todos: Todos,
	}

	tpl.Execute(w, data)
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		//fmt.Fprintf(w, "Parsing Form err: %v", err)
		return
	}

	todoList := ReadJsonFromFile()
	todoList = append(todoList, newTodo(r.FormValue("task"), false))
	//fmt.Fprintf(w, "Task: %s", r.FormValue("task"))
	//printTodos(w, todoList)
	WriteJsonToFile(todoList)
	data := PageData{
		Title: "Your list",
		Todos: todoList,
	}
	tpl.Execute(w, data)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}

	var todoList []Todo
	WriteJsonToFile(todoList)
	data := PageData{
		Title: "Your list",
		Todos: todoList,
	}

	tpl.Execute(w, data)
}

func Create() {
	//var task string
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", todo)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/clear", clearHandler)
	http.ListenAndServe(":"+port, mux)
}
