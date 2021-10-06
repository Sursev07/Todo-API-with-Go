package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"todo-API-with-go/models"
	"encoding/json"
	// "github.com/gorilla/context"
)

func UserAuthorization(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userData := c.MustGet("userData").(jwt.MapClaims)

		if userData != "admin" {
			res := models.ResultErr{Code: 401, Error: "Unauthorized", Message: "you are not authorized to access this data"}
			result, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result)
			return
			return
		}
		nextHandler.ServeHTTP(w, r)
	})
}
