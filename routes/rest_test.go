package routes

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"fmt"
	"github.com/RobinNagpal/auth-service-go/utils"
)

func TestSignup(t *testing.T) {

	incompleteSignup := &signupCommand{
		Name:  "Robin Nagpal",
		Email: "email",
	}
	jsonData, _ := json.Marshal(incompleteSignup)
	fmt.Println(string(jsonData))
	req, _ := http.NewRequest("POST", "/hello/chris", bytes.NewBuffer(jsonData))
	res := httptest.NewRecorder()

	Signup(res, req)

	body := res.Body.String()

	apiError := utils.BadInputError{}
	json.NewDecoder(bytes.NewBuffer([]byte(body))).Decode(&apiError)

	if apiError.Code != "INVALID_INPUT" {
		t.Error("Not Valid Error Code")
	}
	if len(apiError.Details) != 3 {
		t.Error("Not Valid number of field errors")
	}

	emailError := utils.FieldError{
		FieldName: "email",
		Code:      "INVALID_EMAIL",
	}
	if apiError.Details[0] != emailError {
		t.Error("Not valid field email error")
	}

	dobError := utils.FieldError{
		FieldName: "dateOfBirth",
		Code:      "REQUIRED_FIELD",
	}

	if apiError.Details[1] != dobError {
		t.Error("Not valid date of birth field error")
	}

}
