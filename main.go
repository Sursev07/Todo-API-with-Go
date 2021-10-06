package main


import (
	"todo-API-with-go/routes"
	"todo-API-with-go/database"

)

// @title Todo Management API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main()  {
	database.InitDB()
	database.InitialMigration()
	routes.HandleRequest()
}