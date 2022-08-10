package routes

import (
	"GoTask/controllers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.TaskList)
	route.HandleFunc("/add",controllers.TaskAdd).Methods("POST")
	route.HandleFunc("/update",controllers.TaskUpdate).Methods("POST")
	route.HandleFunc("/delete/{id}",controllers.TaskDelete)
	route.HandleFunc("/complete/{id}",controllers.TaskComplete)

	return route
}
