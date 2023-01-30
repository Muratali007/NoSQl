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
	http.HandleFunc("/imageCpu", Handlers.AuthMiddleware(Handlers.ImageCpu))
	http.HandleFunc("/imageCooling", Handlers.AuthMiddleware(Handlers.ImageCooling))
	http.HandleFunc("/imageGpu", Handlers.AuthMiddleware(Handlers.ImageGpu))
	http.HandleFunc("/imageHousing", Handlers.AuthMiddleware(Handlers.ImageHousing))
	http.HandleFunc("/imageMotherboard", Handlers.AuthMiddleware(Handlers.ImageMotherboard))
	http.HandleFunc("/imagePowersupplies", Handlers.AuthMiddleware(Handlers.ImagePowersupplies))
	http.HandleFunc("/images", Handlers.AuthMiddleware(Handlers.Images))
	http.HandleFunc("/register", Handlers.Register)
	http.HandleFunc("/login", Handlers.Login)
	http.HandleFunc("/logout", Handlers.Logout)
	http.HandleFunc("/home", Handlers.AuthMiddleware(Handlers.Home))
	http.HandleFunc("/delete", Handlers.AuthMiddleware(Handlers.DeleteUser))
	http.HandleFunc("/update", Handlers.AuthMiddleware(Handlers.UpdateUser))
	http.HandleFunc("/shop", Handlers.AuthMiddleware(Handlers.Shop))
	http.HandleFunc("/fullListSsd", Handlers.AuthMiddleware(Handlers.ListSsd))
	http.HandleFunc("/fullListCpu", Handlers.AuthMiddleware(Handlers.ListCpu))
	http.HandleFunc("/fullListCooling", Handlers.AuthMiddleware(Handlers.ListCooling))
	http.HandleFunc("/fullListHdd", Handlers.AuthMiddleware(Handlers.ListHdd))
	http.HandleFunc("/fullListHousing", Handlers.AuthMiddleware(Handlers.ListHousing))
	http.HandleFunc("/fullListRam", Handlers.AuthMiddleware(Handlers.ListRam))
	http.HandleFunc("/fullListMotherboard", Handlers.AuthMiddleware(Handlers.ListMotherboard))
	http.HandleFunc("/fullListPowerSupply", Handlers.AuthMiddleware(Handlers.ListPowerSupply))
	http.HandleFunc("/fullListGpu", Handlers.AuthMiddleware(Handlers.ListGpu))
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
