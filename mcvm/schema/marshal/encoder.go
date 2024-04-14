package marshal

import "reflect"

func GetFilledFields(value interface{}) []string {
	filled := []string{}
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	for i := range v.NumField() {
		field := v.Field(i)
		isSet := field.IsValid() && !field.IsZero()
		if isSet {
			filled = append(filled, t.Field(i).Name)
		}
	}

	return filled
}
