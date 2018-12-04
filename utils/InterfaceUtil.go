package utils

import "reflect"

//func IsNil(i interface{}) bool {
//	defer func() {
//		recover()
//	}()
//	vi := reflect.ValueOf(i)
//	return vi.IsNil()
//}

func IsNil(itf interface{}) bool {
	var result bool = false

	if nil == itf {
		result = true
	} else {
		vi := reflect.ValueOf(itf)
		if vi.Kind() == reflect.Ptr {
			result = vi.IsNil()
		}
	}

	return result
}
