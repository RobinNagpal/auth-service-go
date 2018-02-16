package routes

import (
	"net/http"
	"encoding/json"
	"errors"
	"time"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/universal-translator"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	"github.com/go-playground/locales/en"
	"reflect"
	"strings"
)

type signupCommand struct {
	Name        string    `json:"name" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	Password    string    `json:"password" validate:"required,gte=8,lte=30"`
}

type loginCommand struct {
	email    string `validate:"nonzero"`
	password string `validate:"nonzero"`
}

type FieldError struct {
	FieldName string `json:"fieldName"`
	Code     string `json:"code"`
}

type BadRequestError struct {
	Code    string
	Message string
	Details []FieldError
}

func setError(responseWriter http.ResponseWriter, apiError BadRequestError) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)
	byteArray, _ := json.MarshalIndent(apiError, "", "    ")
	fmt.Fprintf(responseWriter, string(byteArray[:]))
}

func parseCommand(w http.ResponseWriter, r *http.Request, c interface{}) error {

	var (
		uni      *ut.UniversalTranslator
		validate *validator.Validate
	)

	validate = validator.New()
	eng := en.New()
	uni = ut.New(eng, eng)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.FindTranslator("en")


	enTranslations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "REQUIRED_FIELD", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})


	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "INVALID_EMAIL", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})


	if r.Body == nil {
		setError(w, BadRequestError{
			Code:    "MISSING_BODY",
			Message: "Body is not present",
		})
		return errors.New("body is bil")
	}
	requestBody := r.Body

	err := json.NewDecoder(requestBody).Decode(c)
	if err != nil {
		setError(w, BadRequestError{
			Code:    "INVALID_JSON",
			Message: "Cant decode the request",
		})
		return errors.New("can't decode")
	}

	fmt.Println(c)
	if validateError := validate.Struct(c); validateError != nil {

		if _, ok := validateError.(*validator.InvalidValidationError); ok {
			setError(w, BadRequestError{
				Code:    "MISSING_BODY",
				Message: "Body is not present",
			})
		}
		// translate all error at once
		var fieldErrors []FieldError
		fmt.Println(validateError)
		validateErrors := validateError.(validator.ValidationErrors)
		for _, err := range validateErrors {
			fieldError := FieldError{
				FieldName: err.Field(),
				Code:     err.Translate(trans),
			}
			fieldErrors = append(fieldErrors, fieldError)
		}

		setError(w, BadRequestError{
			Code:    "INVALID_INPUT",
			Message: "One/Many input fields are invalid",
			Details: fieldErrors,
		})
		return errors.New("Invalid Fields")

	}

	return nil
}
