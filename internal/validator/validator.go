package validator

import (
	"fmt"

	gopgvalidator "github.com/go-playground/validator/v10"
)

func Validate(s any) []string {
	var res []string
	validate := gopgvalidator.New()
	errs := validate.Struct(s)

	if errs != nil {
		errMsgs := make([]string, 0)
		errMap := map[string]string{
			"max":      ": довжина має бути не більше %s символів",
			"gte":      ": значення має бути більше %s",
			"lte":      ": значення має бути менше %s",
			"iscolor":  ": значення має відповідати формату кольора%s",
			"iso4217":  ": значення має відповідати формату iso4217%s",
			"datetime": ": невірний формати дати %s",
		}
		for _, err := range errs.(gopgvalidator.ValidationErrors) {
			errMsgs = append(errMsgs, err.StructField()+fmt.Sprintf(errMap[err.Tag()], err.Param()))
		}
		return errMsgs
	}

	return res
}