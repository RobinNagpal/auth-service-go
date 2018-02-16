package routes

import "net/http"

func Signup(w http.ResponseWriter, r *http.Request) {
	var command signupCommand
	error := validateAndParse(w, r, &command); if error != nil {

	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var command loginCommand
	validateAndParse(w, r, &command)
}

func GetUserInfo(http.ResponseWriter, *http.Request) {

}

func ForgotPassword(http.ResponseWriter, *http.Request) {

}

func ResetPassword(http.ResponseWriter, *http.Request) {

}

func EnableTwoFactAuth(http.ResponseWriter, *http.Request) {

}

func FacebookLogin(http.ResponseWriter, *http.Request) {

}

func GoogleLogin(http.ResponseWriter, *http.Request) {

}

func GetQRCode(http.ResponseWriter, *http.Request) {

}
