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

	ok, err := LoggedUser(r, cfg.Secret, uuid.MustParse(chirp.UserID))
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error checking logged in user: %s", err), http.StatusInternalServerError)
		return
	}
	if ok == false {
		ErrorResponse(w, "Cannot create chirp: invalid authorization", http.StatusBadRequest)
		return
	}

	chirpDB, err := cfg.Database.CreateChirps(context.Background(), database.CreateChirpsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Body:      cleaned_text,
		UserID:    uuid.MustParse(chirp.UserID),
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
		ErrorResponse(w, fmt.Sprintf("GET error request : %s", err), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(chirps)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in marshaling GET request : %s", err), http.StatusBadRequest)
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

func (cfg *ApiConfig) HandlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	chirpID := uuid.MustParse(r.PathValue("chirpsID"))

	chirps, err := cfg.Database.GetChirpID(context.Background(), chirpID)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error getting chirps from chirpID: %s", err), http.StatusNotFound)
		return
	}

	ok, err := LoggedUser(r, cfg.Secret, chirps.UserID)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error checking logged in user: %s", err), http.StatusInternalServerError)
		return
	}
	if ok == false {
		ErrorResponse(w, "Cannot create chirp: invalid authorization", http.StatusBadRequest)
		return
	}

	err = cfg.Database.DeleteChirpsID(context.Background(), chirpID)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error deleting chirp : %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
