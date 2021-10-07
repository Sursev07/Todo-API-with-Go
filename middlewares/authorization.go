package middlewares

import (
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"todo-API-with-go/models"
	"todo-API-with-go/database"
)


func UserAuthorization(nextHandler http.Handler)  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserData := context.Get(r, "userData").(jwt.MapClaims)
		userID := uint(UserData["id"].(float64))
		fmt.Println(UserData)

		params := mux.Vars(r)
		var todo models.Todo
		DB := database.GetDB()
		if err := DB.First(&todo, params["id"]).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(userID, todo.AuthorId,userID == todo.AuthorId, "DATA")
		fmt.Println(todo, "TODOOO")
		if userID != todo.AuthorId {
			res := models.ResultErr{Code: 401, Error: "Unauthorized", Message: "you are not authorized to access this data"}
			result, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result)
			return
		}
		nextHandler.ServeHTTP(w, r)
	})
}
