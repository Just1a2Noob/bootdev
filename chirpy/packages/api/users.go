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
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
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

func (cfg *ApiConfig) HandlerDeleteUsers(w http.ResponseWriter, r *http.Request) {
	err := cfg.Database.DeleteUsers(context.Background())
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in deleting users from database : %s", err), http.StatusInternalServerError)
	}
}
