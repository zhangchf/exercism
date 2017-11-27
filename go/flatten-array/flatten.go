package flatten

import (
	"reflect"
	"fmt"
)

func Flatten(input interface{}) (result []interface{}) {
	flatten(input, result)
	return result
}

func flatten(input interface{}, result []interface{}) {
	value := reflect.ValueOf(input)
	fmt.Println(input, value.Type(), value.Kind())
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			flatten(value.Index(i), result)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64 :
		result = append(result, value.Int())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			fmt.Println(value.Field(i))
		}
	}
}
