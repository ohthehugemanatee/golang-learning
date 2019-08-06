package reflection

import "reflect"

// Walk over an array, executing a function for each string member.
func Walk(x interface{}, userFunc func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	userFunc(field.String())
}
