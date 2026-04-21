package validator

import (
    "errors"
    "fmt"
    "strings"

    "github.com/go-playground/validator/v10"
)

// FieldMessage — custom message per field+tag, opsional
type FieldMessage struct {
    Field   string // nama field, kosongkan untuk match semua field
    Tag     string // nama tag validasi, kosongkan untuk match semua tag
    Message string
}

// ValidateStruct — validasi struct + format error message
// customMessages opsional, fallback ke default message jika tidak ada
func ValidateStruct(v *validator.Validate, s any, customMessages ...FieldMessage) error {
    if err := v.Struct(s); err != nil {
        var validationErrors validator.ValidationErrors
        if errors.As(err, &validationErrors) {
            return fmt.Errorf("%s", formatErrors(validationErrors, customMessages))
        }
        return fmt.Errorf("invalid input")
    }
    return nil
}

func formatErrors(errs validator.ValidationErrors, customMessages []FieldMessage) string {
    var messages []string

    for _, e := range errs {
        msg := findCustomMessage(e, customMessages)
        if msg == "" {
            msg = defaultMessage(e)
        }
        messages = append(messages, msg)
    }

    return strings.Join(messages, ", ")
}

func findCustomMessage(e validator.FieldError, customMessages []FieldMessage) string {
    for _, cm := range customMessages {
        fieldMatch := cm.Field == "" || cm.Field == e.Field()
        tagMatch := cm.Tag == "" || cm.Tag == e.Tag()

        if fieldMatch && tagMatch {
            return cm.Message
        }
    }
    return ""
}

func defaultMessage(e validator.FieldError) string {
    switch e.Tag() {
        case "required":
            return fmt.Sprintf("%s is required or cannot be empty!", e.Field())
        case "email":
            return fmt.Sprintf("%s must be a valid email!", e.Field())
        case "min":
            return fmt.Sprintf("%s must be at least %s characters!", e.Field(), e.Param())
        case "max":
            return fmt.Sprintf("%s must be at most %s characters!", e.Field(), e.Param())
        case "alphaspace", "alphaspaceunicode":
            return fmt.Sprintf("%s cannot contain numbers or symbols!", e.Field())
        case "uuid4":
            return fmt.Sprintf("%s must be an uuid type!", e.Field())
        case "username":
            return fmt.Sprintf("%s must be must be contains (lowercase, number, dot) and cannot contain whitespace!", e.Field())
        case "password":
            return fmt.Sprintf("%s must be contains (1 uppercase, 1 number, 1 symbol) and cannot contain whitespace!", e.Field())
        case "futuredate":
            return fmt.Sprintf("datetime must be now or future!")
        case "oneof":
            return fmt.Sprintf("priority must be %s!", e.Param())
        case "gte":
            return fmt.Sprintf("%s out of range, cannot be less than %s!", e.Field(), e.Param())
        default:
            return fmt.Sprintf("%s is invalid!", e.Field())
    }
}