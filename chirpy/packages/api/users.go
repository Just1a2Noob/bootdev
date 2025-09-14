package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Just1a2Noob/bootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

type EmailResponse struct {
	Email string `json:"email"`
}

func (cfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var emailReq EmailResponse

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&emailReq)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in decoding json request : %s", err), http.StatusNotAcceptable)
	}

	email := emailReq.Email

	user, err := cfg.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email:     email,
	})
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error creating user to database : %v", err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(user)
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
