package handlers

import (
	"as/internal/models"
	"as/pkg"
	"as/utils"
	"net/http"
)

func AddUserPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.Error(err)

	db := pkg.ConnDB.GetConn()
	requiredFields := []string{"name", "lastname", "username", "email", "password"}

	if ValidateRequiredFields(r, requiredFields) {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         "Required fields are missing.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	var user models.User
	user.Name = r.FormValue("name")
	user.LastName = r.FormValue("lastname")

	email := models.Email(r.FormValue("email"))
	msg, valid := email.Validate(db)
	user.Email = email
	if !valid {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         msg,
			Application: "json",
		}
		res.Send(w)
		return
	}

	username := models.Username(r.FormValue("username"))
	msg, valid = username.Validate(db)
	user.Username = username
	if !valid {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         msg,
			Application: "json",
		}
		res.Send(w)
		return
	}

	password := models.Password(r.FormValue("password"))
	hash, err := password.Encrypt()
	utils.Error(err)
	user.Password = models.Password(hash)

	err = db.Create(&user).Error
	if err != nil {
		res := Response{
			Status:      http.StatusInternalServerError,
			Msg:         err.Error(),
			Application: "json",
		}
		res.Send(w)
		return
	}

	res := Response{
		Status:      http.StatusCreated,
		Msg:         "User created successfully.",
		Data:        user,
		Application: "json",
	}
	res.Send(w)
}

func DeleteUserDelete(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	if id == "" {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         "Please provide an id.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	db := pkg.ConnDB.GetConn()

	var user models.User

	if err := db.Where("ID = ?", id).First(&user).Error; err != nil {
		res := Response{
			Status:      http.StatusNotFound,
			Msg:         "User not found.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		res := Response{
			Status:      http.StatusInternalServerError,
			Msg:         err.Error(),
			Application: "json",
		}
		res.Send(w)
		return
	}

	res := Response{
		Status:      http.StatusOK,
		Msg:         "User deleted successfully.",
		Application: "json",
	}
	res.Send(w)
}

func GetUserGet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.Error(err)

	db := pkg.ConnDB.GetConn()

	if ValidateRequiredFields(r, []string{"id"}) {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         "Required fields are missing.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	var user models.User
	id := r.FormValue("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		res := Response{
			Status:      http.StatusNotFound,
			Msg:         "User not found.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	res := Response{
		Status:      http.StatusOK,
		Msg:         "User found.",
		Data:        user,
		Application: "json",
	}
	res.Send(w)
}

func UpdateUserPut(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.Error(err)

	db := pkg.ConnDB.GetConn()

	if ValidateRequiredFields(r, []string{"id"}) {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         "Required fields are missing.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	var user models.User
	id := r.FormValue("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		res := Response{
			Status:      http.StatusNotFound,
			Msg:         "User not found.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	if r.FormValue("name") != "" {
		user.Name = r.FormValue("name")
	}

	if r.FormValue("lastname") != "" {
		user.LastName = r.FormValue("lastname")
	}

	if r.FormValue("email") != "" {
		email := models.Email(r.FormValue("email"))
		msg, valid := email.Validate(db)
		if !valid {
			res := Response{
				Status:      http.StatusBadRequest,
				Msg:         msg,
				Application: "json",
			}
			res.Send(w)
			return
		}
		user.Email = email
	}

	if r.FormValue("username") != "" {
		username := models.Username(r.FormValue("username"))
		msg, valid := username.Validate(db)
		if !valid {
			res := Response{
				Status:      http.StatusBadRequest,
				Msg:         msg,
				Application: "json",
			}
			res.Send(w)
			return
		}
		user.Username = username
	}

	if r.FormValue("password") != "" {
		password := models.Password(r.FormValue("password"))
		hash, err := password.Encrypt()
		utils.Error(err)
		user.Password = models.Password(hash)
	}

	if err := db.Save(&user).Error; err != nil {
		res := Response{
			Status:      http.StatusInternalServerError,
			Msg:         err.Error(),
			Application: "json",
		}
		res.Send(w)
		return
	}

	res := Response{
		Status:      http.StatusOK,
		Msg:         "User updated successfully.",
		Data:        user,
		Application: "json",
	}
	res.Send(w)
}

func AllUsersGet(w http.ResponseWriter, r *http.Request) {
	db := pkg.ConnDB.GetConn()

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		res := Response{
			Status:      http.StatusInternalServerError,
			Msg:         err.Error(),
			Application: "json",
		}
		res.Send(w)
		return
	}

	res := Response{
		Status:      http.StatusOK,
		Msg:         "Users found.",
		Data:        users,
		Application: "json",
	}
	res.Send(w)
}
