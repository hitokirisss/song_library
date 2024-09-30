package routes

import (
	"database/sql"

	"song_library/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/songs", handlers.GetSongs(db)).Methods("GET")
	r.HandleFunc("/songs/{id}/text", handlers.GetSongText).Methods("GET")
	r.HandleFunc("/songs", handlers.AddSong).Methods("POST")
	r.HandleFunc("/songs/{id}", handlers.UpdateSong).Methods("PUT")
	r.HandleFunc("/songs/{id}", handlers.DeleteSong).Methods("DELETE")
	return r
}
