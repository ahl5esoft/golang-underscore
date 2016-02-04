package underscore

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func ParseJson(str string, container interface{}) error {
	reader := strings.NewReader(str)
	return json.NewDecoder(reader).Decode(container)
}

func ToJson(value interface{}) (string, error) {
	var err error
	res := ""

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
		case reflect.String:
			res = value.(string)
			break
		case reflect.Array, 
				reflect.Map,
				reflect.Slice,
				reflect.Struct:
			var bytes []uint8
			bytes, err = json.Marshal(value)
			if err == nil {
				res = string(bytes)
			}
			break
		case reflect.Bool:
			res = strconv.FormatBool(value.(bool))
			break
		case reflect.Float32, reflect.Float64:
			res = strconv.FormatFloat(
				rv.Float(),
				'f', 
				-1, 
				64,
			)
			break
		case reflect.Int,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64,
				reflect.Int8:
			res = strconv.FormatInt(
				rv.Int(),
				10,
			)
			break
		case reflect.Ptr:
			res, err = ToJson(
				reflect.Indirect(rv).Interface(),
			)
			break
		case reflect.Uint,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64,
				reflect.Uint8:
			res = strconv.FormatUint(
				rv.Uint(),
				10,
			)
			break
	}
	return res, err
}

func ToRealValue(rv reflect.Value) interface{} {
	var value interface{}
	switch rv.Kind() {
		case reflect.Bool:
			value = rv.Bool()
			break
		case reflect.Float32, reflect.Float64:
			value = rv.Float()
			break
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			value = rv.Int()
			break
		case reflect.String:
			value = rv.String()
			break
		case reflect.Struct:
			value = rv.Interface()
			break
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = rv.Uint()
			break
		case reflect.Ptr:
			return ToRealValue(
				reflect.Indirect(rv),
			)
		default:
			if !rv.IsNil() {
				value = rv.Interface()
			}
			break
	}
	return value
}

/*
	@source		数据源,array or map
	@iterator	迭代器
		@in_1	source的元素
		@in_2	下标或者map的key
		@out_1	提前终止标志
		@out_2	error(reflect.Value)
*/
func each(source interface{}, iterator func([]reflect.Value) (bool, reflect.Value)) error {
	if source == nil {
		return nil
	}

	sourceRV := reflect.ValueOf(source)
	switch sourceRV.Kind() {
		case reflect.Array:
		case reflect.Slice:
			if sourceRV.Len() == 0 {
				return nil
			}

			for i := 0; i < sourceRV.Len(); i++ {
				isBreak, errRV := iterator(
					[]reflect.Value{
						sourceRV.Index(i),
						reflect.ValueOf(i),
					},
				)
				
				if isErrorRVValid(errRV) {
					return errRV.Interface().(error)
				}

				if isBreak {
					break
				}
			}
		case reflect.Map:
			keyRVs := sourceRV.MapKeys()
			if len(keyRVs) == 0 {
				return nil
			}

			for i := 0; i < len(keyRVs); i++ {
				isBreak, errRV := iterator(
					[]reflect.Value{
						sourceRV.MapIndex(keyRVs[i]),
						keyRVs[i],
					},
				)
				
				if isErrorRVValid(errRV) {
					return errRV.Interface().(error)
				}

				if isBreak {
					break
				}
			}
	}
	return nil
}

func getPropertyRV(entityRV reflect.Value, property string) (reflect.Value, error) {
	var err error
	rv := entityRV.FieldByName(property)
	if !rv.IsValid() {
		err = errors.New("invalid field: [" + property + "]")
	}

	return rv, err
}

func getPropertyValue(entity interface{}, property string) (interface{}, error) {
	rv, err := getPropertyRV(
		reflect.ValueOf(entity),
		property,
	)
	if err == nil {
		return rv.Interface(), nil
	}
	
	return nil, err
}

func isErrorRVValid(errRV reflect.Value) bool {
	return errRV.IsValid() && !errRV.IsNil()
}

func makeSliceRV(rt reflect.Type, length int) reflect.Value {
	return reflect.MakeSlice(rt, length, length)
}

func makeSliceRVWithElem(elemRT reflect.Type, length int) reflect.Value {
	sliceType := reflect.SliceOf(elemRT)
	return makeSliceRV(sliceType, length)
}

func makeMapRV(keyRT, valueRT reflect.Type) reflect.Value {
	mapRT := reflect.MapOf(keyRT, valueRT)
	return reflect.MakeMap(mapRT)
}

func makeGroupRV(keyRT, elemRT reflect.Type) reflect.Value {
	sliceRT := reflect.SliceOf(elemRT)
	return makeMapRV(keyRT, sliceRT)
}