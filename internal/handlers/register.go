package handlers

import (
	"as/pkg/schemas"
	"as/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Msg         string      `json:"msg"`
	Data        interface{} `json:"data"`
	Application string
}

func (r *Response) Marshal() []byte {
	/* Only Status, Msg, Data*/
	jsonUser := map[string]interface{}{
		"status": r.Status,
		"msg":    r.Msg,
		"data":   r.Data,
	}

	/* Marshal */
	res, err := json.Marshal(jsonUser)
	utils.Error(err)
	return res
}

func (r *Response) SendRes(w http.ResponseWriter) {
	w.Header().Set("Content-Type", fmt.Sprintf("application/%s", r.Application))
	w.WriteHeader(r.Status)
	_, err := w.Write(r.Marshal())
	utils.Error(err)
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {

	/*
		| TASK LIST |
		2. Check if the username and email already exists
		3. If not, create a new user with JWT token and save it to the database
		4. Redirect to the Home page
	*/

	/* request body */

	err := r.ParseForm()
	utils.Error(err)

	/* check if the username and email already exists */
	valid := schemas.Find("username", r.FormValue("username"))
	if valid {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         "Username already exists",
			Application: "json",
		}
		res.SendRes(w)
		return
	}

	valid = schemas.Find("email", r.FormValue("email"))
	if valid {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         "Email already exists",
			Application: "json",
		}
		res.SendRes(w)
		return
	}

	/* create a new user and save it to the database */
	user := schemas.User{
		Name:     r.FormValue("name"),
		LastName: r.FormValue("lastname"),
		Username: r.FormValue("username"),
		Email:    schemas.Email{Value: r.FormValue("email")},
		Password: schemas.Password{Value: r.FormValue("password")},
	}

	/* validate email */
	valid, err = user.Email.Validate()
	utils.Error(err)
	if valid == false {
		res := Response{
			Status:      http.StatusBadRequest,
			Msg:         "Invalid email",
			Application: "json",
		}
		res.SendRes(w)
		return
	}

	/* encrypt password */
	user.Password.Value, err = user.Password.Encrypt()
	utils.Error(err)

	err = user.Insert()
	utils.Error(err)

	/* redirect to the Home page */
	res := Response{
		Status:      http.StatusOK,
		Msg:         "User created successfully",
		Application: "json",
	}
	res.SendRes(w)
}
