package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// val := reflect.ValueOf(x)
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value
	switch val.Kind() {
	case reflect.Struct:
		// walk through all the fields in struct
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		// walk through all the elements in slice
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.String:
		// reach what we want, call fn()
		numberOfValues = 0
		fn(val.String())
	}

	// if val is a slice, it will not get into this loop
	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
