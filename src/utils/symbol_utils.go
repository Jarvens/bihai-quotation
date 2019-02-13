// auth: kunlun
// date: 2019-02-12
// description:
package utils

import (
	"errors"
	"reflect"
)

// compare obj & target
// validate target is contains obj
// return true or false and error
func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}

	}
	return false, errors.New("not in array")
}
