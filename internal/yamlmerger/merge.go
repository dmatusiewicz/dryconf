package yamlmerger

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/dmatusiewicz/dryconf/internal/logs"
)

var (
	errTypeMissmatch = errors.New("Type missmatch between dst and src")
	errDSTnotPointer = errors.New("dst should be a pointer")
)
var ret reflect.Value
var logger logs.Log

// Merge merge two structures into one.
func Merge(dst, src map[string]interface{}) (interface{}, error) {
	return merge(dst, src)
}

func merge(dst, src map[string]interface{}) (interface{}, error) {

	dstV := reflect.ValueOf(dst)
	srcV := reflect.ValueOf(src)
	ret = reflect.MakeMap(dstV.Type())
	val, err := deepMerge(dstV, srcV)
	return val.Interface(), err
}

func deepMerge(dst, src reflect.Value) (reflect.Value, error) {
	var err error
	switch src.Kind() {
	case reflect.Map:
		for _, key := range dst.MapKeys() {
			ret.SetMapIndex(key, dst.MapIndex(key))
		}
		for _, key := range src.MapKeys() {
			if ret.MapIndex(key).IsValid() {
				switch reflect.TypeOf(ret.MapIndex(key).Interface()).Kind() {
				case reflect.String:
					ret.SetMapIndex(key, src.MapIndex(key))
				case reflect.Slice:
					retSlice := reflect.ValueOf(dst.MapIndex(key).Interface())
					srcSlice := reflect.ValueOf(src.MapIndex(key).Interface())
					dstSlice := reflect.MakeSlice(retSlice.Type(), 0, retSlice.Len())
					dstSlice = reflect.AppendSlice(retSlice, srcSlice)
					// fmt.Println(dstSlice)
					ret.SetMapIndex(key, dstSlice)
				case reflect.Map:
					fmt.Println("Map case")
				}

			} else {
				ret.SetMapIndex(key, src.MapIndex(key))
			}
		}
	}
	return ret, err
}
