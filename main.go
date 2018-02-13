package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/auth/signup", signup).Methods("POST")
	router.HandleFunc("/auth/login", login).Methods("POST")
	router.HandleFunc("/auth/user", getUserInfo).Methods("GET")
	router.HandleFunc("/auth/forgot-password", forgotPassword).Methods("POST")
	router.HandleFunc("/auth/reset-password", resetPassword).Methods("POST")
	router.HandleFunc("/auth/enableTwoFactAuth", enableTwoFactAuth).Methods("PUT")
	router.HandleFunc("/auth/facebook-login", facebookLogin).Methods("POST")
	router.HandleFunc("/auth/google-login", googleLogin).Methods("POST")
	router.HandleFunc("/auth/qr-code", getQRCode).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(http.ResponseWriter, *http.Request) {

}

func login(http.ResponseWriter, *http.Request) {

}

func getUserInfo(http.ResponseWriter, *http.Request) {

}


func forgotPassword(http.ResponseWriter, *http.Request) {

}

func resetPassword(http.ResponseWriter, *http.Request) {

}


func enableTwoFactAuth(http.ResponseWriter, *http.Request) {

}

func facebookLogin(http.ResponseWriter, *http.Request) {

}

func googleLogin(http.ResponseWriter, *http.Request) {

}

func getQRCode(http.ResponseWriter, *http.Request) {

}
