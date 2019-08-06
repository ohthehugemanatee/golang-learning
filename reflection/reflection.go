package reflection

import "reflect"

// Walk over an array, executing a function for each string member.
func Walk(x interface{}, userFunc func(input string)) {
	val := getValue(x)
	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		userFunc(val.String())
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	}

	for i := 0; i < numberOfValues; i++ {
		Walk(getField(i).Interface(), userFunc)
	}
}

func getValue(x interface{}) (val reflect.Value) {
	val = reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return
}
