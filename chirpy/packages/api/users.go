package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Just1a2Noob/bootdev/chirpy/internal/auth"
	"github.com/Just1a2Noob/bootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

type UserRequest struct {
	Password           string `json:"password"`
	Email              string `json:"email"`
	Expires_in_seconds int    `json:"expire_in_seconds"`
}

const DefaultExpires = 3600

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
}

func (cfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userReq UserRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&userReq)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in decoding json request : %s", err), http.StatusNotAcceptable)
		return
	}

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

	user, err := cfg.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Email:          email,
		HashedPassword: hashed,
	})

	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error creating user to database : %v", err), http.StatusInternalServerError)
		return
	}

	userRes := UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	data, err := json.Marshal(userRes)
	if err != nil {
		ErrorResponse(w, "Error encoding user entry to json", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (cfg *ApiConfig) HandlerLoginUser(w http.ResponseWriter, r *http.Request) {
	var userReq UserRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&userReq)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in decoding json request : %s", err), http.StatusNotAcceptable)
		return
	}

	user, err := cfg.Database.SearchUser(context.Background(), userReq.Email)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Cannot find email in database : %s", err), http.StatusNotFound)
		return
	}

	err = auth.CheckPasswordHash(userReq.Password, user.HashedPassword)
	if err != nil {
		ErrorResponse(w, "Password is wrong, please try again", http.StatusNotFound)
	}

	if userReq.Expires_in_seconds < 0 {
		ErrorResponse(w, "ExpiresAt cannot be negative", http.StatusNotAcceptable)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.Secret, time.Duration(userReq.Expires_in_seconds))
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error creating JWT token: %s", err), http.StatusInternalServerError)
		return
	}

	_, err = auth.ValidateJWT(token, cfg.Secret)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Token missmatch with secret: %s", err), http.StatusInternalServerError)
	}

	userRes := UserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     token,
	}

	data, err := json.Marshal(userRes)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error marshalling response : %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (cfg *ApiConfig) HandlerDeleteUsers(w http.ResponseWriter, r *http.Request) {
	err := cfg.Database.DeleteUsers(context.Background())
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in deleting users from database : %s", err), http.StatusInternalServerError)
	}
}
