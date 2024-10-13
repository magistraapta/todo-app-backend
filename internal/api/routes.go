package api

import (
	"net/http"
	"todo-app-backend/internal/api/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/task", handlers.GetAllTask)
}
