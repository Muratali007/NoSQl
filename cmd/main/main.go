package main

import (
	"MongoDb/Handlers"
	"MongoDb/internal/data"
	"MongoDb/pkg/logging"
	"net/http"
)

func main() {
	logger := logging.GetLogger()
	data.Init("test", "users")
	logger.Info("Create route")
	http.HandleFunc("/register", Handlers.Register)
	http.HandleFunc("/login", Handlers.Login)
	http.HandleFunc("/logout", Handlers.Logout)
	http.HandleFunc("/home", Handlers.AuthMiddleware(Handlers.Home))
	http.HandleFunc("/shop", Handlers.AuthMiddleware(Handlers.Shop))
	http.HandleFunc("/list", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/list2", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/list3", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/list4", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/list5", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/list6", Handlers.AuthMiddleware(Handlers.List))
	http.ListenAndServe(":8080", nil)
}
