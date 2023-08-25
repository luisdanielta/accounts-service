package handlers

import (
	"as/internal/models"
	"as/pkg/schemas"
	"as/utils"
	"encoding/json"
	"net/http"
)

/* type response, status, msg, data */

func RegisterPost(w http.ResponseWriter, r *http.Request) {

	/*
		| TASK LIST |
		1. Geting information from the form, name, username, email, password
		2. Check if the username and email already exists
		3. If not, create a new user with JWT token and save it to the database
		4. Redirect to the Home page
	*/

	/* request body */
	err := r.ParseForm()
	utils.HttpError(err, w)

	user := models.User{
		Name:     r.FormValue("name"),
		LastName: r.FormValue("lastname"),
		Username: r.FormValue("username"),
		Email: models.Email{
			Value: r.FormValue("email"),
		},
		Password: models.Password{
			Value: r.FormValue("password"),
		},
	}

	/* encrypt password */
	hash, err := user.Password.Encrypt()
	utils.HttpError(err, w)
	user.Password.Value = hash

	/* check out Email */
	valid, err := user.Email.Validate()
	utils.HttpError(err, w)

	if !valid {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	/* add user - database */

	utils.HttpError(err, w)

	/* response */
	jsonUser := map[string]interface{}{
		"status": 201,
		"msg":    "User created successfully",
		"data":   user,
	}

	dbUser := schemas.User{}

	res, err := json.Marshal(jsonUser)
	utils.HttpError(err, w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
