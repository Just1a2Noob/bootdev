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

func (cfg *ApiConfig) HandlerAddChirp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var chirpReq chirpRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&chirpReq)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in decoding json request : %s", err), http.StatusNotAcceptable)
	}

	cleaned_text := chirpReq.Body

	chirp, err := cfg.Database.CreateChirps(context.Background(), database.CreateChirpsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Body:      cleaned_text,
		UserID:    uuid.MustParse(chirpReq.User),
	})
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error creating chirp for database : %s", err), http.StatusInternalServerError)
	}

	data, err := json.Marshal(chirp)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error converting chirp entry to json : %s", err), http.StatusInternalServerError)
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

func (cfg *ApiConfig) HandlerGetChirps(w http.ResponseWriter, r http.Request) {
	chirps, err := cfg.Database.GetChirps(context.Background())
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("GET error request : %s", err), http.StatusInternalServerError)
	}

	data, err := json.Marshal(chirps)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in marshaling GET request : %s", err), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
