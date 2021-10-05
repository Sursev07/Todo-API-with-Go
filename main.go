package main


import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"todo-API-with-go/controllers"
	"todo-API-with-go/database"
	"net/http"
)

func main()  {
	database.InitDB()
	database.InitialMigration()
	r := mux.NewRouter()


	r.HandleFunc("/", controllers.WelcomeAPI).Methods("GET")
	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos/:id", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/:id", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/:id", controllers.DeleteTodo).Methods("DELETE")

	fmt.Println("Server is starting at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}