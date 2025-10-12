package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Just1a2Noob/bootdev/chirpy/internal/database"
	"github.com/Just1a2Noob/bootdev/chirpy/packages/api"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	fmt.Printf("DB_URL: %s\n", os.Getenv("DB_URL"))
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database : %s", err)
	}

	dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"
	secret := os.Getenv("SECRET")
	polka_key := os.Getenv("POLKA_KEY")

	apicfg := api.ApiConfig{
		FileserverHits: atomic.Int32{},
		Database:       *dbQueries,
		Secret:         secret,
		Polka_key:      polka_key,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", apicfg.MiddlewareMetricsInc(
		http.FileServer(http.Dir(filepathRoot)),
	)))

	mux.HandleFunc("GET /admin/metrics", apicfg.HandlerMetrics)

	mux.HandleFunc("POST /admin/reset", apicfg.HandlerDeleteUsers)
	mux.HandleFunc("DELETE /admin/users", apicfg.HandlerDeleteUser)
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", apicfg.HandlerValidate)
	mux.HandleFunc("POST /api/users", apicfg.HandlerCreateUser)
	mux.HandleFunc("POST /api/chirps", apicfg.HandlerAddChirp)
	mux.HandleFunc("GET /api/chirps", apicfg.HandlerGetChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apicfg.HandlerGetChirpID)
	mux.HandleFunc("POST /api/login", apicfg.HandlerLoginUser)
	mux.HandleFunc("POST /api/refresh", apicfg.HandlerRefreshToken)
	mux.HandleFunc("POST /api/revoke", apicfg.HandlerRevokeToken)
	mux.HandleFunc("PUT /api/users", apicfg.HandlerUpdateUser)
	mux.HandleFunc("DELETE /api/chirps/{chirpsID}", apicfg.HandlerDeleteChirp)

	mux.HandleFunc("POST /api/polka/webhooks", apicfg.HandlerChirpRed)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
