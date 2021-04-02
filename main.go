package main

import (
	"log"
	"net/http"
	userHandler "totality-corp-assignment/handlers/user"
	userService "totality-corp-assignment/services/user"
	userStore "totality-corp-assignment/stores/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	store := userStore.New()
	service := userService.New(store)
	handler := userHandler.New(service)

	r.HandleFunc("/user/{id}", handler.GetByID)
	r.HandleFunc("/users/{ids}", handler.GetByIDs)
	log.Fatal("Stopped ", http.ListenAndServe(":8001", r))
}
