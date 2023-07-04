package controllers

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"

	"github.com/AlexSparr0w/mmrapi/model"
	"github.com/AlexSparr0w/mmrapi/utils"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var validate *validator.Validate

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var input model.CreateUserRequest

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Body should be in a valid format")
		return
	}

	userID := uuid.New()

	user := &model.User{
		ID:       userID.String(),
		Email:    input.Email,
		Password: input.Password,
	}

	model.DB.Create(user)

	apikey := RandStringBytes(12)

	apikeyModel := &model.Apikey{
		ID:     uuid.New().String(),
		UserID: userID.String(),
		Apikey: apikey,
	}

	model.DB.Create(apikeyModel)

	userResposne := &model.UserResponse{
		Email:  input.Email,
		Apikey: apikey,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResposne)
}

func CompleteRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	apikey := r.Header.Get("apikey")

	check := model.DB.Exec("SELECT * FROM apikeys WHERE apikey = $1", apikey)
	if check.RowsAffected == 0 {
		utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	var input model.CompleteRegistrationReuqest
	var userID string

	userIDcheck := model.DB.Raw("SELECT user_id FROM apikeys WHERE apikey = $1", apikey).Scan(&userID)
	if userIDcheck.RowsAffected == 0 {
		utils.RespondWithError(w, http.StatusUnprocessableEntity, "no user found")
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Body should be in a valid format")
		return
	} else {
		model.DB.Exec("UPDATE users SET total_matches = $1, total_mmr = $2 WHERE id = $3", input.TotalMatches, input.TotalMmr, userID)
	}

}
