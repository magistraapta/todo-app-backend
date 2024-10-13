package handlers

import (
	"fmt"
	"net/http"
)

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting all task...")
}
