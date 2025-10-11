package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/Just1a2Noob/bootdev/chirpy/internal/database"
	"github.com/google/uuid"
)

// Config holds application configuration
type ApiConfig struct {
	FileserverHits atomic.Int32
	Database       database.Queries
	Secret         string
}

func (cfg *ApiConfig) HandlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
<html>

<body>
	<h1>Welcome, Chirpy Admin</h1>
	<p>Chirpy has been visited %d times!</p>
</body>

</html>
	`, cfg.FileserverHits.Load())))
}

func (cfg *ApiConfig) MiddlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.FileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *ApiConfig) HandlerResetHits(w http.ResponseWriter, r *http.Request) {
	cfg.FileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}

func SuccessResponse(w http.ResponseWriter, req *chirpRequest, cfg *ApiConfig) {
	cleaned_text := ProfaneChirp(req.Body)
	response := chirpRequest{
		Body: cleaned_text,
		User: req.User,
	}

	data, err := json.Marshal(response)
	if err != nil {
		ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Saves chirpReq in chirps database
	_, err = cfg.Database.CreateChirps(context.Background(), database.CreateChirpsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Body:      req.Body,
		UserID:    uuid.MustParse(req.User),
	})
	if err != nil {
		ErrorResponse(w, "Error inserting chirp to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ErrorResponse(w http.ResponseWriter, message string, code int) {
	response := APIResponse{Error: message}

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling error response : %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "internal server error"}`))
	}

	w.WriteHeader(code)
	w.Write(data)
}

// TODO: Create a function to handle user logging check.
// input(userID) -> checks loggedUser == userID -> output(boolean)
