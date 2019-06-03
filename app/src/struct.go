package main

type registerPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type user struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Firstname    string `json:"firstname"`
	Middlename   string `json:"middlename"`
	Lastname     string `json:"lastname"`
	Organization string `json:"organization"`
	Logins       int    `json:"logins"` //count the number if logins
}

type authToken struct {
	Token string `json:"token"`
}

type userResponse struct {
	*user
	authToken
}

type errorResponse struct {
	Error string `json:"error"`
}
