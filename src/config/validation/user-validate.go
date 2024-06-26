package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	toolkit "github.com/renatofagalde/golang-toolkit"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *toolkit.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	var rest_err toolkit.RestErr

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type") //json enviando errado
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Cause{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Cause{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}

}
