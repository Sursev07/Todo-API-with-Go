package routes


import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"todo-API-with-go/controllers"
)

func HandleRequest()  {
	r := mux.NewRouter()


	r.HandleFunc("/", controllers.WelcomeAPI).Methods("GET")
	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	fmt.Println("Server is starting at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}