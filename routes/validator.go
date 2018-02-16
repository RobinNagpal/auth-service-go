package routes

import (
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/locales/en"
	"reflect"
	"strings"
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
)

var (
	uni        *ut.UniversalTranslator
	validate   *validator.Validate
	translator *ut.Translator
)

func init() {
	validate = validator.New()
	eng := en.New()
	uni = ut.New(eng, eng)

	trans, _ := uni.FindTranslator("en")
	translator = &trans
	regCustomErrorMsgs(validate, translator)
	regJsonNames(validate)
}

func validateAndParse(w http.ResponseWriter, r *http.Request, c interface{}) error {

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

	if validateError := validate.Struct(c); validateError != nil {


		if _, ok := validateError.(*validator.InvalidValidationError); ok {
			setError(w, BadRequestError{
				Code:    "MISSING_BODY",
				Message: "Body is not present",
			})
		}
		// translate all error at once
		var fieldErrors []FieldError
		validateErrors := validateError.(validator.ValidationErrors)
		for _, err := range validateErrors {
			fieldError := FieldError{
				FieldName: err.Field(),
				Code:      err.Translate(*translator),
			}
			fieldErrors = append(fieldErrors, fieldError)
		}

		setError(w, BadRequestError{
			Code:    "INVALID_INPUT",
			Message: "One/Many input fields are invalid",
			Details: fieldErrors,
		})
		return errors.New("invalid fields")

	}

	return nil
}

func setError(responseWriter http.ResponseWriter, apiError BadRequestError) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)
	byteArray, _ := json.MarshalIndent(apiError, "", "    ")
	fmt.Fprintf(responseWriter, string(byteArray[:]))
}

func regJsonNames(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func regCustomErrorMsgs(validate *validator.Validate, translator *ut.Translator) {
	validate.RegisterTranslation("required", *translator, func(ut ut.Translator) error {
		return ut.Add("required", "REQUIRED_FIELD", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterTranslation("email", *translator, func(ut ut.Translator) error {
		return ut.Add("email", "INVALID_EMAIL", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
}
