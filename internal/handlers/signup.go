package handlers

import "net/http"

func SignupGet(w http.ResponseWriter, r *http.Request) {

	/*
		| TASK LIST |
		1. Geting information from the form, name, username, email, password
		2. Check if the username and email already exists
		3. If not, create a new user with JWT token and save it to the database
		4. Redirect to the Home page
	*/
}
