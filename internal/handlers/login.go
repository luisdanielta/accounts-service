package handlers

import (
	"as/internal/models"
	"as/pkg"
	"as/pkg/env"
	"as/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = env.JWT_KEY

func generateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	utils.Error(err)

	var user models.User

	db := pkg.ConnDB.GetConn()

	// Get username and password from database
	if err := db.Where("username = ?", r.FormValue("username")).First(&user).Error; err != nil {
		res := Response{
			Status:      http.StatusNotFound,
			Msg:         "Invalid username or password.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	// Compare password
	if err := user.Password.Compare(r.FormValue("password")); err != nil {
		res := Response{
			Status:      http.StatusNotFound,
			Msg:         "Invalid username or password.",
			Application: "json",
		}
		res.Send(w)
		return
	}

	// You can generate and send authentication token here if needed
	token, err := generateToken(user)
	utils.Error(err)

	// Add the token to the response headers for authorization
	w.Header().Set("Authorization", "Bearer "+token)

	res := Response{
		Status:      http.StatusOK,
		Msg:         "Login successful.",
		Data:        user,
		Application: "json",
	}
	res.Send(w)
}
