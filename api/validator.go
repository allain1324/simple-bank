package api

import (
	"github.com/allain1324/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validConcurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)

	}
	return false
}
