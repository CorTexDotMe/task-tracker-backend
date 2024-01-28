package resolver

import (
	"encoding/json"
	"net/http"
	"task-tracker-backend/internal/model"
	"task-tracker-backend/internal/repository/gorm"
	"task-tracker-backend/internal/utils"
)

// Return jwt token if request body has correct User credentials in form-data format
func Login(w http.ResponseWriter, r *http.Request) {
	input := model.Credentials{Username: r.FormValue("username"), Password: r.FormValue("password")}
	authenticated := gorm.GetUserRepository().Authenticate(input)

	if !authenticated {
		http.Error(w, "wrong username or password", http.StatusNotFound)
		return
	}

	token, err := utils.GenerateToken(input.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(token)
}

// Create new user with credentials from request body in form-data format.
//
// Return created user as json
func Register(w http.ResponseWriter, r *http.Request) {
	input := model.Credentials{Username: r.FormValue("username"), Password: r.FormValue("password")}
	user, err := gorm.GetUserRepository().SaveFromInput(model.CredsToNewUser(input))
	utils.HandleError(err) // TODO dont panic

	json.NewEncoder(w).Encode(user)
}
