package reflection

import "reflect"

// Walk over an array, executing a function for each string member.
func Walk(x interface{}, userFunc func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			userFunc(field.String())
		}
		if field.Kind() == reflect.Struct {
			Walk(field.Interface(), userFunc)
		}
	}
}
