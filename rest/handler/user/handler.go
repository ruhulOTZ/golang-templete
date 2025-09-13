package user

import (
	"encoding/json"
	"expenseTracker/config"
	"expenseTracker/database"
	"expenseTracker/utils"
	"fmt"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.GetUsers(), http.StatusOK)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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
		utils.SendError(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := utils.CreateJWT(cnf.JwtSecretKey, utils.Payload{
		Sub:       user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})

	if err != nil {
		utils.SendError(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, accessToken, http.StatusOK)
}
