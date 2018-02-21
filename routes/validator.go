package routes

import (
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"github.com/RobinNagpal/auth-service-go/utils"
)


func validateAndParse(w http.ResponseWriter, r *http.Request, c interface{}) error {

	if r.Body == nil {
		setError(w, utils.BadInputError{
			Code:    "MISSING_BODY",
			Message: "Body is not present",
		})
		return errors.New("body is bil")
	}
	requestBody := r.Body

	err := json.NewDecoder(requestBody).Decode(c)

	if err != nil {
		setError(w, utils.BadInputError{
			Code:    "INVALID_JSON",
			Message: "Cant decode the request",
		})
		return errors.New("can't decode")
	}

	if commandError := utils.ValidateStruct(c); commandError != nil {
		setError(w, *commandError)
		return errors.New("invalid request")
	}

	return nil
}

func setError(responseWriter http.ResponseWriter, apiError utils.BadInputError) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)
	byteArray, _ := json.MarshalIndent(apiError, "", "    ")
	fmt.Fprintf(responseWriter, string(byteArray[:]))
}

