package handler

import (
	"encoding/json"
	"expenseTracker/database"
	"expenseTracker/utils"
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.GetUsers(), http.StatusOK)
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", http.StatusBadRequest)
		return
	}

	response := database.AddUser(&newUser)

	utils.SendData(w, response, 201)
}

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var ReqLogin = &ReqLogin{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(ReqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid json", http.StatusBadRequest)
		return
	}

	user := database.FindUser(ReqLogin.Email, ReqLogin.Password)

	if user == nil {
		utils.SendError(w, "Invalid email or password",  http.StatusUnauthorized)
		return
	}

	utils.SendData(w, user, http.StatusOK)
}