package controllers


import (
	"fmt"
	"todo-API-with-go/models"
	"net/http"
	"encoding/json"
	"io/ioutil"

	// "gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"github.com/gorilla/mux"
	"todo-API-with-go/database"
)

// var DB *gorm.DB
var err error


// GetTodos godoc
// @Summary Get details of all todos
// @Description Get details of all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func GetTodos(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var todos []models.Todo
	DB := database.GetDB()
	DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
	
}

func WelcomeAPI(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	result := "Hello Welcome to Todo Management API"
	json.NewEncoder(w).Encode(result)
	
}


// GetTodo godoc
// @Summary Get details of todo
// @Description Get details of todo
// @Tags todo
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {array} models.Todo
// @Router /todos/{id}  [get]
func GetTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo models.Todo
	DB := database.GetDB()
	data := DB.First(&todo, params["id"])
	res := models.Result{Code: 200, Data: todo, Message: "Success Get Todo"}
	result, err := json.Marshal(res)

	fmt.Println("masukk", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	
}


// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo with the input paylod
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body Todo true "Todo has been created"
// @Success 201 {object} models.Todo
// @Router /todos [post]
func CreateTodo(w http.ResponseWriter, r *http.Request)  {
	payloads, _ := ioutil.ReadAll(r.Body)

	var todo models.Todo
	json.Unmarshal(payloads, &todo)
	DB := database.GetDB()
	DB.Create(&todo)

	res := models.Result{Code: 201, Data: todo, Message: "Todo has been created"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update a todo with the input paylod
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body Todo true "Todo has been created"
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request)  {
	fmt.Print("Masukkkk")
	vars := mux.Vars(r)
	todoID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var todoEdit models.Todo
	json.Unmarshal(payloads, &todoEdit)

	var todo models.Todo
	DB := database.GetDB()
	DB.First(&todo, todoID)
	DB.Model(&todo).Updates(todoEdit)

	 res := models.Result{Code: 200, Data: todo, Message: "Todo has been updated"}
	 result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
	
}


// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo with the input paylod
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200
// @Router /todos/{id}[delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var todo models.Todo
	DB := database.GetDB()
	DB.Delete(&todo, params["id"])
	res := models.Result{Code: 200, Message: "Todo has been deleted"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}