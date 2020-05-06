package validator

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Setup() {
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic(errors.New("Binding Validator Engine Error\n"))
	}

	err := validate.RegisterValidation("mgtv_url", MgtvUrl)
	if err != nil {
		panic(err)
	}
}
