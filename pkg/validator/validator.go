package validator

import (
	"fmt"
	validatorMain "github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"strings"
)

type (
	ValidateSchema struct {
		Validator *validatorMain.Validate
	}
	Message struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func New() *ValidateSchema {
	validate := &ValidateSchema{Validator: validatorMain.New()}
	validate.Validator.RegisterTagNameFunc(validate.getFieldName)

	return validate
}

func (v *ValidateSchema) getFieldName(fld reflect.StructField) string {
	filed := fld.Tag.Get("json")
	name := strings.SplitN(filed, ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}

func (v *ValidateSchema) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

func (v *ValidateSchema) ValidateHandleError(i interface{}) []Message {
	if err := v.Validate(i); err != nil {
		var messages []Message
		for _, err := range err.(validatorMain.ValidationErrors) {
			field := v.namespace(err.Namespace(), err.Namespace())
			messages = append(messages, Message{
				Field:   field,
				Message: v.matchMessage(err, field),
			})
		}

		return messages
	}

	return nil
}

func (v *ValidateSchema) namespace(key, message string) string {
	splitField := strings.Split(key, ".")
	if len(splitField) > 1 {
		field := ""
		for _, value := range strings.Split(key, ".") {
			if regexp.MustCompile("^[a-z]").MatchString(value) {
				if field == "" {
					field = value
				} else {
					field += "." + value
				}
			}
		}

		message = strings.Replace(message, key, field, 1)
	}

	return message
}

func (v *ValidateSchema) matchMessage(err validatorMain.FieldError, field string) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "oneof":
		values := strings.Split(err.Param(), " ")
		return fmt.Sprintf("This field one of value: %s", strings.Join(values, ", "))
	case "numeric":
		return "This field may only contain numeric characters"
	default:
		return v.namespace(err.Namespace(), err.Error())
	}
}
