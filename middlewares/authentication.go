package middlewares

import (
	"todo-API-with-go/helpers"
	"todo-API-with-go/models"
	"net/http"
	"encoding/json"
	"github.com/gorilla/context"
)

func Authentication(nextHandler http.Handler)  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		verifyToken, err := helpers.VerifyToken(r)
		_ = verifyToken

		w.Header().Set("Content-Type", "application/json")
		if err != nil {

			res := models.ResultErr{Code: 401, Error: "Unauthenticated", Message: err.Error()}
			result, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(result)
			return
		}
		
		context.Set(r, "userData", verifyToken)
        nextHandler.ServeHTTP(w, r)
	})
}
