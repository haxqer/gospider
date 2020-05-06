package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var validMacRaw = regexp.MustCompile(`^https?://www\.mgtv\.com/\w+/\d+/\d+\.html$`)

func MgtvUrl(fl validator.FieldLevel) bool {
	if validMacRaw.MatchString(fl.Field().String()) {
		return true
	}

	return false
}
