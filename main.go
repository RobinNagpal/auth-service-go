package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/RobinNagpal/auth-service-go/routes"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/auth/signup", routes.Signup).Methods("POST")
	router.HandleFunc("/auth/login", routes.Login).Methods("POST")
	router.HandleFunc("/auth/user", routes.GetUserInfo).Methods("GET")
	router.HandleFunc("/auth/forgot-password", routes.ForgotPassword).Methods("POST")
	router.HandleFunc("/auth/reset-password", routes.ResetPassword).Methods("POST")
	router.HandleFunc("/auth/enableTwoFactAuth", routes.EnableTwoFactAuth).Methods("PUT")
	router.HandleFunc("/auth/facebook-login", routes.FacebookLogin).Methods("POST")
	router.HandleFunc("/auth/google-login", routes.GoogleLogin).Methods("POST")
	router.HandleFunc("/auth/qr-code", routes.GetQRCode).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}