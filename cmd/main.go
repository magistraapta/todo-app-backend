package main

import (
	"fmt"
	"todo-app-backend/internal/api"
)

func main() {
	api.SetupRoutes()
	fmt.Println("hello world")
}
