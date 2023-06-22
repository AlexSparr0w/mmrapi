package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/AlexSparr0w/mmrapi/model"
	"github.com/AlexSparr0w/mmrapi/utils"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input CreateUserRequest

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Body should be in a valid format")
		return
	}

	user := &model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	model.DB.Create(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
