package http

import (
	"encoding/json"
	"errors"
	"github/m4nsur/todo-api-lrn/todo"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
		respondWithError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	if err := taskDTO.ValidateForCreate(); err != nil {
		respondWithError(w, http.StatusBadRequest, "validation failed: "+err.Error())
		return
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			respondWithError(w, http.StatusConflict, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "failed to create task")
		}
		return
	}

	respondWithJSON(w, http.StatusCreated, todoTask)
}

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if title == "" {
		respondWithError(w, http.StatusBadRequest, "task title is required")
		return
	}

	task, err := h.todoList.GetTask(title)
	if err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "failed to get task")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *HTTPHandlers) HandleGetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *HTTPHandlers) HandleGetAllUncompletedTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListNotCompletedTasks()
	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if title == "" {
		respondWithError(w, http.StatusBadRequest, "task title is required")
		return
	}

	task, err := h.todoList.CompleteTask(title); 
	if err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "failed to complete task")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if title == "" {
		respondWithError(w, http.StatusBadRequest, "task title is required")
		return
	}

	if err := h.todoList.DeleteTask(title); err != nil {
		if errors.Is(err, todo.ErrTaskNotFound) {
			respondWithError(w, http.StatusNotFound, err.Error())
		} else {
			respondWithError(w, http.StatusInternalServerError, "failed to delete task")
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	errDTO := ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errDTO)
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"message":"failed to encode response"}`, http.StatusInternalServerError)
	}
}