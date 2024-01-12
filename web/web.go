package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
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

func WriteJsonToFile(data []Todo) {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

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

	WriteJsonToFile(todoList)
	data := PageData{
		Title: "Your list",
		Todos: todoList,
	}
	tpl.Execute(w, data)
}

func remove(s []Todo, i int) []Todo {
	s = append(s[:i], s[i+1:]...)
	return s
}

func toggleHandler(w http.ResponseWriter, r *http.Request) {
	todos := ReadJsonFromFile()
	var index struct {
		Index int `json:"index"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&index); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	todos[index.Index].Done = !todos[index.Index].Done

	jsonResponse(w, todos)
	WriteJsonToFile(todos)
}

// Helper function to send JSON responses
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
func clearDoneHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		return
	}

	todoList := ReadJsonFromFile()

	for i := 0; i < len(todoList); i++ {
		if todoList[i].Done == true {
			todoList = remove(todoList, i)
		}
	}

	WriteJsonToFile(todoList)
	data := PageData{
		Title: "Your list",
		Todos: todoList,
	}

	tpl.Execute(w, data)
}

func clearAllHandler(w http.ResponseWriter, r *http.Request) {
	var todoList []Todo
	WriteJsonToFile(todoList)
	data := PageData{
		Title: "Your list",
		Todos: todoList,
	}

	tpl.Execute(w, data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Serving on port", port, "...")

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", todo)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/clearDone", clearDoneHandler)
	mux.HandleFunc("/clearAll", clearAllHandler)
	mux.HandleFunc("/toggle", toggleHandler)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
