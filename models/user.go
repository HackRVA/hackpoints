package models

import "github.com/dgrijalva/jwt-go"

// swagger:response userResponse
type userResponse struct {
	Body Member
}

// tokenResponseBody for json response of signin
// swagger:response loginResponse
type tokenResponseBody struct {
	// in: body
	Body TokenResponse
}

// swagger:parameters loginRequest
type loginRequest struct {
	// in: body
	Body Credentials
}

// swagger:parameters registerUserRequest
type userRegisterRequest struct {
	// in: body
	Body Credentials
}

// Credentials Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	// Password - the user's password
	// required: true
	// example: string
	Password string `json:"password"`
	// Email - the users email
	// required: true
	// example: string
	Email string `json:"email"`
}

// Member -- a member of the makerspace
type Member struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserResponse - a user object that we can send as json
type UserResponse struct {
	// Email - user's Email
	// example: string
	Email string `json:"email"`
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// TokenResponse -- for json response of signin
type TokenResponse struct {
	// login response to send token string
	//
	// Example: <TOKEN_STRING>
	Token string `json:"token"`
}
