package transform

import (
	"errors"
	"reflect"
	"strings"
)

func T(s any) error {
	fieldType := reflect.TypeOf(s).Elem()
	fieldValue := reflect.ValueOf(s).Elem()

	if fieldValue.Kind() != reflect.Struct {
		return errors.New("Expected a struct")
	}

	//	percorrer todos os campos da struct
	for i := 0; i < fieldType.NumField(); i++ {
		field := fieldType.Field(i)
		value := fieldValue.Field(i)
		if !value.CanSet() {
			continue
		}

		transform := field.Tag.Get("transform")
		if transform == "" || transform == "-" {
			continue
		}

		switch value.Kind() {
		case reflect.String:
			strValue := value.String()
			if strings.Contains(transform, "upper") {
				value.SetString(strings.ToUpper(strValue))

			}
			if strings.Contains(transform, "lower") {
				value.SetString(strings.ToLower(strValue))
			}
		default:
			return errors.New("unsupported type")
		}
	}
	return nil
}
