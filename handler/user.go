package handler

import (
	"fmt"
	"net/http"
	"user/db"
	"user/model"
	"user/utils"

	"github.com/google/uuid"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	database, err := db.DbIn()
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to connect to database: %v", err), http.StatusInternalServerError)
		return
	}
	defer database.Close()

	db.CreateUserTable(database)

	var newUser model.User
	if err := utils.ParseJson(r, &newUser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateUser(newUser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	id, err := db.InsertUser(database, newUser)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, map[string]uuid.UUID{"id": id})
}
