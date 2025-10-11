package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Just1a2Noob/bootdev/chirpy/internal/auth"
	"github.com/Just1a2Noob/bootdev/chirpy/internal/database"
)

func (cfg *ApiConfig) HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	// Decodes user request
	w.Header().Set("Content-Type", "application/json")

	var userReq UserRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&userReq)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in decoding json request : %s", err), http.StatusNotAcceptable)
		return
	}

	// Gets the header token
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error getting authorization header: %s", err), http.StatusNotFound)
		return
	}

	// Validates token
	// tokenUserID can be assumed as a unique identifier for the user
	tokenUserID, err := auth.ValidateJWT(token, cfg.Secret)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Failed validating JWT token: %s", err), http.StatusBadRequest)
		return
	}

	// Search the user based on token user ID
	_, err = cfg.Database.SearchUserWithID(context.Background(), tokenUserID)
	if err != nil {
		ErrorResponse(w, "Error token provided does not have a matching id", http.StatusNotFound)
		return
	}

	// Given the new password hash it
	email := userReq.Email
	hashed, err := auth.HashPassword(userReq.Password)
	if err != nil {
		ErrorResponse(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	err = auth.CheckPasswordHash(userReq.Password, hashed)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Password and hashed mismatched : %s", err), http.StatusInternalServerError)
	}

	// Updates the users table
	err = cfg.Database.UpdateUser(context.Background(), database.UpdateUserParams{
		ID:             tokenUserID,
		Email:          email,
		HashedPassword: hashed,
		UpdatedAt:      time.Now().UTC(),
	})
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error updating entry : %s", err), http.StatusInternalServerError)
		return
	}

	userRes := UserResponse{
		ID:        tokenUserID,
		Email:     email,
		UpdatedAt: time.Now().UTC(),
	}

	data, err := json.Marshal(userRes)
	if err != nil {
		ErrorResponse(w, "Error encoding user entry to json", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
