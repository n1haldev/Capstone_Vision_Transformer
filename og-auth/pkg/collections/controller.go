package collections

import (
	"encoding/json"
	
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/naren/og-auth/pkg/models"
	"github.com/naren/og-auth/pkg/utils"
)

var NewReg models.Reg

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.Reg{}
	utils.ParseBody(r, CreateUser)

	// Check if the email already exists in the database
	existingUser := models.GetUserByEmail(CreateUser.Email)
	if existingUser != nil {
		http.Error(w, "Email already exists", http.StatusConflict)
		return
	}

	// Hash the password before creating the user
	hashedPassword, err := utils.HashPassword(CreateUser.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	CreateUser.Password = hashedPassword

	reg := CreateUser.CreateUser()

	res, err := json.Marshal(reg)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	allUsers := models.GetAllUsers()

	res, err := json.Marshal(allUsers)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	// Set the response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CheckUser(w http.ResponseWriter, r *http.Request) {
    CreateUser := &models.Reg{}
    utils.ParseBody(r, CreateUser)
    ans := models.GetUserByEmailLogin(CreateUser.Email, CreateUser.Password)
    if ans == "success" {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "login successful"})
        return
    }

    if ans == "fail" {
        http.Error(w, `{"error": "email/password is incorrect"}`, http.StatusUnauthorized)
        return
    }
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user := models.DeleteUser(ID)
	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
