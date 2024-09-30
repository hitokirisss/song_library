package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"song_library/database"
	"song_library/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger" // Импортируем для Swagger
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.MigrateDB(db)

	r := routes.SetupRoutes(db)

	// Подключаем Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Println("Server is running on port :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
