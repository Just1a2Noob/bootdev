package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const MaxChirpLength = 140

var profanities = []string{"kerfuffle", "sharbert", "fornax"}

// Request and Response types
type chirpRequest struct {
	Body string `json:"body"`
}

type APIResponse struct {
	Valid bool   `json:"valid,omitempty"`
	Error string `json:"error,omitempty"`
	Body  string `json:"body"`
}

// Custom error Handling

type ValidationError struct {
	Message string
	Code    int
}

func (e ValidationError) Error() string {
	return e.Message
}

func HandlerValidate(w http.ResponseWriter, r *http.Request) {
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

	SuccessResponse(w, chirpReq)
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

func SuccessResponse(w http.ResponseWriter, req *chirpRequest) {
	cleaned_text := ProfaneChirp(req.Body)
	response := APIResponse{Body: cleaned_text}

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error in marshaling valid response : %s", err)
		ErrorResponse(w, "Internal Server Error", http.StatusInternalServerError)
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
