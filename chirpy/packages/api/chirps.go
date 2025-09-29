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

type Chirp struct {
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}

func (cfg *ApiConfig) HandlerAddChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var chirp Chirp

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&chirp)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in decoding json request : %s", err), http.StatusNotAcceptable)
		return
	}

	cleaned_text := chirp.Body

	// Gets authentication/login user
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error getting authorization header: %s", err), http.StatusInternalServerError)
		return
	}

	tokenUserID, err := auth.ValidateJWT(token, cfg.Secret)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Failed validating JWT token: %s", err), http.StatusBadRequest)
		return
	}

	// Parse and validate the user UUID
	userID, err := uuid.Parse(chirp.UserID)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Invalid user ID format: %s", err), http.StatusBadRequest)
		return
	}

	// Checking if login is the same as userID from chirp
	if userID != tokenUserID {
		ErrorResponse(w, fmt.Sprintf("Chirp user ID does not match with Logged in user: %s", err), http.StatusUnauthorized)
		return
	}

	chirpDB, err := cfg.Database.CreateChirps(context.Background(), database.CreateChirpsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Body:      cleaned_text,
		UserID:    userID,
	})
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error creating chirp for database : %s", err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(chirpDB)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error converting chirp entry to json : %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (cfg *ApiConfig) HandlerDeleteChirps(w http.ResponseWriter, r http.Request) {
	err := cfg.Database.DeleteChirps(context.Background())
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in deleting chirps from database : %s", err), http.StatusInternalServerError)
	}
}

func (cfg *ApiConfig) HandlerGetChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.Database.GetChirps(context.Background())
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("GET error request : %s", err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(chirps)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in marshaling GET request : %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (cfg *ApiConfig) HandlerGetChirpID(w http.ResponseWriter, r *http.Request) {
	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error parsing chirpID : %s", err), http.StatusInternalServerError)
		return
	}

	chirpDB, err := cfg.Database.GetChirpID(context.Background(), chirpID)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Chirp ID did not match with database : %s", err), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(chirpDB)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error marshalling chirp to json : %s", err), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
