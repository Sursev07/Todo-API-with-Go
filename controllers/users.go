package controllers

import (
	// "fmt"
	"todo-API-with-go/models"
	"todo-API-with-go/helpers"
	"net/http"
	"encoding/json"
	"io/ioutil"
	// "github.com/gorilla/mux"
	"todo-API-with-go/database"
)

func UserRegister(w http.ResponseWriter, r *http.Request)  {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(payloads, &user)
	DB := database.GetDB()
	DB.Create(&user)

	res := models.Result{Code: 200, Data: user, Message: "Todo has been created"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UserLogin(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	payloads, _ := ioutil.ReadAll(r.Body)

	var user models.User
	json.Unmarshal(payloads, &user)
	DB := database.GetDB()

	password := ""

	password = user.Password

	err := DB.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		res := models.ResultErr{Code: 400, Error: "Unauthorized", Message: "Invalid email/password"}
		result, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(result)
		// c.JSON(http.StatusUnauthorized, gin.H{
		// 	"error":   "Unauthorized",
		// 	"message": "invalid email/password",
		// })
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		res := models.ResultErr{Code: 400, Error: "Unauthorized", Message: "Invalid email/password"}
		result, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(result)
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	// c.JSON(http.StatusOK, gin.H{
	// 	"token": token,
	// })

	res := models.Result{Code: 200, Data: token, Message: "Successfully login"}
	result, err := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}