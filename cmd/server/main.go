package main

import (
	"golang-ddd/internal/app"
	"golang-ddd/internal/infra/repo"
	ifs "golang-ddd/internal/interfaces/http"
	"golang-ddd/pkg/db"
	"log"
	"net/http"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	userRepo := repo.NewUserRepo(db)
	pUserService := app.NewUserService(userRepo)
	userHandler := ifs.UserHandler{PUserService: pUserService}

	mux := http.NewServeMux()
	userHandler.RegisterRoutes(mux)

	port := ":8080"
	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
