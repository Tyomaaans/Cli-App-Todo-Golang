package validator

import (
    "regexp"
    "unicode"
    "time"

    "github.com/go-playground/validator/v10"
)

var instance *validator.Validate

func New() *validator.Validate {
    instance = validator.New()
    registerCustomValidators(instance)
    return instance
}

func registerCustomValidators(v *validator.Validate) {
    v.RegisterValidation("alphaspaceunicode", alphaSpaceUnicode)
    v.RegisterValidation("password", validatePassword)
    v.RegisterValidation("username", validateUsername)
    v.RegisterValidation("futuredate", validateFutureDate)
}

func validatePassword(fl validator.FieldLevel) bool {
    s := fl.Field().String()

    var hasUpper, hasNumber, hasSymbol bool

    for _, c := range s {
        switch {
        case unicode.IsUpper(c):
            hasUpper = true
        case unicode.IsDigit(c):
            hasNumber = true
        case !unicode.IsLetter(c) && !unicode.IsDigit(c) && !unicode.IsSpace(c):
            hasSymbol = true
        case unicode.IsSpace(c):
            return false 
        }
    }

    return hasUpper && hasNumber && hasSymbol
}

func validateUsername(fl validator.FieldLevel) bool {
    regex := regexp.MustCompile(`^[a-z0-9._]+$`)
    return regex.MatchString(fl.Field().String())
}

func alphaSpaceUnicode(fl validator.FieldLevel) bool {
    for _, r := range fl.Field().String() {
        if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
            return false
        }
    }
    return true
}

func validateFutureDate(fl validator.FieldLevel) bool {
    dateStr, ok := fl.Field().Interface().(string)
    if !ok {
        return false
    }

    date, err := time.Parse("2006-01-02", dateStr)
    if err != nil {
        return false
    }

    today := time.Now().Truncate(24 * time.Hour)
    dateOnly := date.Truncate(24 * time.Hour)

    return !dateOnly.Before(today)
}