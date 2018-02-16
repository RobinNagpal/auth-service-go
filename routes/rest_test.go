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
	fmt.Println("#############", res.Body, "#############",)
	if body != "Hello, world" {
		t.Error("Fail! It should not use the default, it should see Chris!")
	}
}
