package reflection

import "reflect"

// Walk over an array, executing a function for each string member.
func Walk(x interface{}, userFunc func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), userFunc)
		}
	case reflect.String:
		userFunc(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			Walk(val.Field(i).Interface(), userFunc)
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
