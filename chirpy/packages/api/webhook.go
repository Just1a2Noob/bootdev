package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Just1a2Noob/bootdev/chirpy/internal/auth"
	"github.com/Just1a2Noob/bootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

type WebHook struct {
	Event string `json:"event"`
	Data  struct {
		UserID string `json:"user_id"`
	} `json:"data"`
}

func (cfg *ApiConfig) HandlerChirpRed(w http.ResponseWriter, r *http.Request) {
	var webHook WebHook

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&webHook)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error in decoding json request : %s", err), http.StatusNotAcceptable)
		return
	}

	if webHook.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	polka_key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Polka key cannot be extracted: %s", err), http.StatusBadRequest)
		return
	}

	if polka_key != cfg.Polka_key {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userID := uuid.MustParse(webHook.Data.UserID)

	_, err = cfg.Database.SearchUserWithID(context.Background(), userID)
	if err != nil {
		ErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}

	err = cfg.Database.SetChirpyRed(context.Background(), database.SetChirpyRedParams{
		ID:          userID,
		IsChirpyRed: "true",
	})

	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error updating chirpy red: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}
