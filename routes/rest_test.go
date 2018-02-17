package routes

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"bytes"
	"fmt"
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

	apiError := BadRequestError{}
	json.NewDecoder(bytes.NewBuffer([]byte(body))).Decode(&apiError)

	if apiError.Code != "INVALID_INPUT" {
		t.Error("Not Valid Error Code")
	}
	if len(apiError.Details) != 3 {
		t.Error("Not Valid number of field errors")
	}

	emailError := FieldError{
		FieldName: "email",
		Code:      "INVALID_EMAIL",
	}
	if apiError.Details[0] != emailError {
		t.Error("Not valid field email error")
	}

	dobError := FieldError{
		FieldName: "dateOfBirth",
		Code:      "REQUIRED_FIELD",
	}

	if apiError.Details[1] != dobError {
		t.Error("Not valid date of birth field error")
	}

}
