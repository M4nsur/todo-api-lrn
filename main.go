package main

import (
	"fmt"
	"github/m4nsur/todo-api-lrn/http"
	"github/m4nsur/todo-api-lrn/todo"
)


func main() {
	todoList := todo.NewList
	httpHandlers := http.NewHttpHandlers(todoList())
	httpServer := http.NewHTTPServer(httpHandlers)
	if err := httpServer.StartServer(); err != nil {

		fmt.Println("failed start server", err)
	}
}