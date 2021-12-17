package errs

import (
	"fmt"
	"strings"
)

func throwServerErrorWithFields(errKey ServerErrKey, fieldValues ...string) error {
	e := buildServerError(errKey)
	if len(fieldValues) == 0 {
		return e
	}

	// order and quantity of fieldValues is guaranteed
	var fieldToValues string = ""
	fields := strings.Split(e.fields, ",")
	for i, field := range fields {
		fieldToValues += fmt.Sprintf("%s=%s", field, fieldValues[i])
		if i < (len(fields) - 1) {
			fieldToValues += ", "
		}
	}
	e.fieldToValues = &fieldToValues
	return e
}

func ThrowServerErrorInvalidValue(className string, fieldName string, value string) error {
	return throwServerErrorWithFields(ErrInvalidValue, className, fieldName, value)
}

func ThrowServerErrorNoValue(class string, field string) error {
	return throwServerErrorWithFields(ErrNoValue, class, field)
}
