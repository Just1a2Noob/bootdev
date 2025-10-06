package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Just1a2Noob/bootdev/chirpy/internal/auth"
	"github.com/Just1a2Noob/bootdev/chirpy/internal/database"
)

func (cfg *ApiConfig) HandlerRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Gets the refresh token
	header := r.Header
	old_token, err := auth.GetBearerToken(header)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error getting bearer token: %s", err), http.StatusInternalServerError)
		return
	}

	// Get refresh token database
	refresh_token, err := cfg.Database.GetUserFromRefreshToken(context.Background(), old_token)
	if err != nil {
		ErrorResponse(w, "Error getting refresh token", http.StatusBadRequest)
		return
	}

	// Checks if user is revoked
	if refresh_token.RevokedAt.Valid {
		ErrorResponse(w, "Refresh token has been revoked", http.StatusBadRequest)
		return
	}

	// Creates tokens with the specified duration
	new_token, err := auth.MakeJWT(refresh_token.UserID, cfg.Secret, time.Duration(60*time.Minute))
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error creating JWT token: %s", err), http.StatusInternalServerError)
		return
	}

	// Validates JWT token
	_, err = auth.ValidateJWT(new_token, cfg.Secret)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Token missmatch with secret: %s", err), http.StatusInternalServerError)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: new_token,
	}

	data, err := json.Marshal(response)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error marshalling response : %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (cfg *ApiConfig) HandlerRevokeToken(w http.ResponseWriter, r *http.Request) {
	header := r.Header

	old_token, err := auth.GetBearerToken(header)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Error getting bearer token: %s", err), http.StatusInternalServerError)
		return
	}

	// Validates the refresh token
	user, err := cfg.Database.GetUserFromRefreshToken(context.Background(), old_token)
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Token is not validated: %s", err), http.StatusBadRequest)
		return
	}

	if user.RevokedAt.Valid {
		ErrorResponse(w, "User is revoked from refreshing token", http.StatusBadRequest)
		return
	}

	// Set User revoke status and updated_at
	var nt sql.NullTime
	nt.Time = time.Now().UTC()
	nt.Valid = true
	err = cfg.Database.RevokeUserToken(context.Background(), database.RevokeUserTokenParams{
		UserID:    user.UserID,
		RevokedAt: nt,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		ErrorResponse(w, fmt.Sprintf("Cannot revoke user : %s", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
