package main

import (
	"MongoDb/Handlers"
	"MongoDb/pkg/logging"
	"net/http"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("Create route")
	http.HandleFunc("/register", Handlers.Register)
	http.HandleFunc("/login", Handlers.Login)
	http.HandleFunc("/logout", Handlers.Logout)
	http.HandleFunc("/home", Handlers.AuthMiddleware(Handlers.Home))
	http.HandleFunc("/shop", Handlers.AuthMiddleware(Handlers.Shop))
	http.ListenAndServe(":8080", nil)
}
