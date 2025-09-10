package api

import (
	"context"
	"encoding/json"
	"log"
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
		log.Fatalf("Error decoding email : %v", err)
		return
	}

	user, err := cfg.Database.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email:     emailReq.Email,
	})
	if err != nil {
		log.Fatalf("Error creating email to database : %s", err)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error encoding email to json : %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
