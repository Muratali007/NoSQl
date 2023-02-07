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
	http.HandleFunc("/listCpu", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/listCooling", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/listHdd", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/listHousing", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/listRam", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/listMotherboard", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/listSsd", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/listPowersupply", Handlers.AuthMiddleware(Handlers.List))
	http.HandleFunc("/listGpu", Handlers.AuthMiddleware(Handlers.List))
	http.ListenAndServe(":8080", nil)
}
