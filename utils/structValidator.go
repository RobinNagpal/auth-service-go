package utils

import (
	"github.com/go-playground/universal-translator"
	"reflect"
	"strings"
	"gopkg.in/go-playground/validator.v9"
	"github.com/go-playground/locales/en"
)

type FieldError struct {
	FieldName string `json:"fieldName"`
	Code      string `json:"code"`
}

type BadInputError struct {
	Code    string
	Message string
	Details []FieldError
}

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

func ValidateStruct(s interface{}) *BadInputError {
	if validateError := validate.Struct(s); validateError != nil {

		if _, ok := validateError.(*validator.InvalidValidationError); ok {
			return &BadInputError{
				Code:    "MISSING_BODY",
				Message: "Body is not present",
			}
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

		return &BadInputError{
			Code:    "INVALID_INPUT",
			Message: "One/Many input fields are invalid",
			Details: fieldErrors,
		}
	}
	return nil
}
