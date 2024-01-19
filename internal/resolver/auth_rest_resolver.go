package resolver

import (
	"encoding/json"
	"net/http"
	"task-tracker-backend/internal/model"
	"task-tracker-backend/internal/repository"
	"task-tracker-backend/internal/utils"
)

var userRepository = repository.UserRepository{}

func Login(w http.ResponseWriter, r *http.Request) {
	input := model.Credentials{Username: r.FormValue("username"), Password: r.FormValue("password")}
	authenticated := userRepository.Authenticate(input)

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
func Register(w http.ResponseWriter, r *http.Request) {
	input := model.Credentials{Username: r.FormValue("username"), Password: r.FormValue("password")}
	user, err := userRepository.SaveFromInput(model.CredsToNewUser(input))
	utils.HandleError(err) // TODO dont panic

	token, err := utils.GenerateToken(user.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(token)
}
