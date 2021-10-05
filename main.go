package main


import (
	"todo-API-with-go/routes"
	"todo-API-with-go/database"

)

func main()  {
	database.InitDB()
	database.InitialMigration()
	routes.HandleRequest()
}