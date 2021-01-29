package tool

import (
	"reflect"
	"strings"
)

/**
空判断：仅仅支持string、slice
*/
func IsEmpty(obj interface{}) bool {
	targetValue := reflect.ValueOf(obj)
	//fmt.Println("===>", targetValue, reflect.TypeOf(obj).Kind())
	switch reflect.TypeOf(obj).Kind() {
	case reflect.String:
		if len(strings.TrimSpace(obj.(string))) == 0 {
			return true
		}
	case reflect.Slice:
		if targetValue.Len() <= 0 {
			return true
		}
	}
	return false
}

/**
判断elem是否在collection中，collection支持的类型 array、slice、map
*/
func Contain(collection interface{}, elem interface{}) bool {
	targetValue := reflect.ValueOf(collection)
	switch reflect.TypeOf(collection).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == elem {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(elem)).IsValid() {
			return true
		}
	}

	return false
}
