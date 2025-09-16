package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

const MaxChirpLength = 140

var profanities = []string{"kerfuffle", "sharbert", "fornax"}

// Request and Response types
type chirpRequest struct {
	Body string `json:"body,omitempty"`
	User string `json:"user"`
}

type APIResponse struct {
	Error string `json:"error,omitempty"`
}

// Custom error Handling

type ValidationError struct {
	Message string
	Code    int
}

func (e ValidationError) Error() string {
	return e.Message
}

func (cfg *ApiConfig) HandlerValidate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	chirpReq, err := parseRequest(r)
	if err != nil {
		if validErr, ok := err.(ValidationError); ok {
			ErrorResponse(w, validErr.Message, validErr.Code)
		}
		return
	}

	// Validate chirp content
	if err := ValidateChirp(chirpReq.Body); err != nil {
		if validationErr, ok := err.(ValidationError); ok {
			ErrorResponse(w, validationErr.Message, validationErr.Code)
		} else {
			ErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	SuccessResponse(w, chirpReq, cfg)
}

func parseRequest(r *http.Request) (*chirpRequest, error) {
	var chirpReq chirpRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&chirpReq); err != nil {
		return nil, ValidationError{
			Message: "Invalid JSON format",
			Code:    http.StatusBadRequest,
		}
	}

	return &chirpReq, nil
}

func ValidateChirp(body string) error {
	trimmedBody := strings.TrimSpace(body)

	if len(trimmedBody) == 0 {
		return ValidationError{
			Message: "Body string cannot be empty",
			Code:    http.StatusBadRequest,
		}
	}

	if len(body) > MaxChirpLength {
		return ValidationError{
			Message: "Chirp is too long",
			Code:    400,
		}
	}

	return nil
}

func ProfaneChirp(body string) string {
	str_arr := strings.Split(body, " ")

	for i, str := range str_arr {
		lower_str := strings.ToLower(str)
		for _, profane := range profanities {
			if strings.Contains(lower_str, profane) {
				str_arr[i] = "****"
				break
			}
		}
	}

	return strings.Join(str_arr, " ")
}
