package http

import (
	"encoding/json"
	"github/m4nsur/todo-api-lrn/todo"
	"net/http"
	"time"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHttpHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return 
	}
}

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	
}

func (h *HTTPHandlers) HandleGetAllTask(w http.ResponseWriter, r *http.Request) {
	
}

func (h *HTTPHandlers) HandleGetAllUncompletedTask(w http.ResponseWriter, r *http.Request) {
	
}

func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	
}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
}