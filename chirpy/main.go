package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/Just1a2Noob/bootdev/chirpy/internal/database"
	"github.com/Just1a2Noob/bootdev/chirpy/packages/api"
	_ "github.com/lib/pq"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database : %s", err)
	}

	dbQueries := database.New(db)

	const filepathRoot = "."
	const port = "8080"

	apicfg := api.ApiConfig{
		FileserverHits: atomic.Int32{},
		Database:       *dbQueries,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", apicfg.MiddlewareMetricsInc(
		http.FileServer(http.Dir(filepathRoot)),
	)))

	mux.HandleFunc("GET /admin/metrics", apicfg.HandlerMetrics)
	mux.HandleFunc("POST /admin/reset", apicfg.HandlerResetHits)
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", api.HandlerValidate)
	mux.HandleFunc("POST /api/users", apicfg.HandlerCreateUser)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
