package reflection

import "reflect"

// Walk over an array, executing a function for each string member.
func Walk(x interface{}, userFunc func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), userFunc)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			userFunc(field.String())
		case reflect.Struct:
			Walk(field.Interface(), userFunc)
		}
	}
}

func getValue(x interface{}) (val reflect.Value) {
	val = reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return
}
