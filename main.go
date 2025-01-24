package main

import (
	"log"
	"net/http"

	"github.com/RomainMarcazzan/ApiRest/config"
	"github.com/RomainMarcazzan/ApiRest/handlers"
	"github.com/RomainMarcazzan/ApiRest/repositories"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	repositories.InitDB()

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/users", handlers.HandleUsers)
	mux.HandleFunc("/v1/users/", handlers.HandleUserByID)
	mux.HandleFunc("/v1/notif", handlers.HandleNotifs)
	mux.HandleFunc("/v1/notif/", handlers.HandleNotifByID)
	mux.HandleFunc("/v1/user/pro-positions", handlers.HandleProPosition)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
