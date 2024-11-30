package handlers

import (
	"dlls/contracts"
	"encoding/json"
	"net/http"
)

func NewAuthHandler(
	authService contracts.AuthService,
) contracts.AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

type authHandler struct {
	authService contracts.AuthService
}

// Login implements contracts.AuthHandler.
func (a *authHandler) Login(w http.ResponseWriter, r *http.Request) {

	var request loginRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := a.authService.Login(request.Email, request.Password)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}

// SignUp implements contracts.AuthHandler.
func (a *authHandler) SignUp(w http.ResponseWriter, r *http.Request) {

	var request signUpRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.authService.SignUp(request.Email, request.Password)

	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
	})
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": message,
		"status":  status,
	})
}
