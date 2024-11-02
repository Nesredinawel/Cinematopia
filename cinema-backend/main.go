package main

import (
	"cinema-backend/handlers" // Your custom handlers
	"database/sql"
	"log"
	"net/http"

	gorillahandlers "github.com/gorilla/handlers" // Alias the import to avoid name conflicts
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	// Auth routes
	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")

	// Movie routes
	router.HandleFunc("/movies", handlers.GetMoviesHandler).Methods("GET")
	router.HandleFunc("/movies", handlers.CreateMovieHandler).Methods("POST")

	// Director routes
	router.HandleFunc("/directors", handlers.GetDirectorsHandler).Methods("GET")
	router.HandleFunc("/directors", handlers.CreateDirectorHandler).Methods("POST")

	// Star routes
	router.HandleFunc("/stars", handlers.GetStarsHandler).Methods("GET")
	router.HandleFunc("/stars", handlers.CreateStarHandler).Methods("POST")

	// Movie Stars routes
	router.HandleFunc("/movie-stars", handlers.GetMovieStarsHandler).Methods("GET")
	router.HandleFunc("/movie-stars", handlers.CreateMovieStarHandler).Methods("POST")

	// Image routes
	// router.HandleFunc("/images", handlers.GetImagesHandler).Methods("GET")
	router.HandleFunc("/images", handlers.CreateImageHandler).Methods("POST")
	router.HandleFunc("/images", handlers.GetImagesHandler).Methods("GET")
	// Bookmark routes
	router.HandleFunc("/bookmarks", handlers.GetUserBookmarksHandler).Methods("GET")
	router.HandleFunc("/bookmarks", handlers.CreateBookmarkHandler).Methods("POST")

	router.HandleFunc("/schedules", handlers.GetSchedulesHandler).Methods("GET")
	router.HandleFunc("/schedules", handlers.CreateScheduleHandler).Methods("POST")

	router.HandleFunc("/seats", handlers.GetSeatsHandler).Methods("GET")
	router.HandleFunc("/seats", handlers.CreateSeatHandler).Methods("POST")

	router.HandleFunc("/tickets", handlers.GetTicketsHandler).Methods("GET")
	router.HandleFunc("/tickets", handlers.CreateTicketHandler).Methods("POST")
	// Configure CORS using gorilla/handlers
	corsHandler := gorillahandlers.CORS(
		gorillahandlers.AllowedOrigins([]string{"http://localhost:3000"}),                             // Allow your frontend origin
		gorillahandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), // Allowed headers
		gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),           // Allowed HTTP methods
	)

	// Start the server with CORS-enabled routes
	log.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", corsHandler(router)))
}
