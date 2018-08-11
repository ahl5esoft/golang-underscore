package underscore

import (
	"errors"
	"reflect"
	"strings"
)

// PropertyRV is 获取reflect.Value
func PropertyRV(name string) func(interface{}) (reflect.Value, error) {
	return func(item interface{}) (reflect.Value, error) {
		var itemRV reflect.Value
		if reflect.TypeOf(item) == NullRvOfRt {
			itemRV = item.(reflect.Value)
		} else {
			itemRV = reflect.ValueOf(item)
		}

		itemRT := itemRV.Type()
		for i := 0; i < itemRT.NumField(); i++ {
			if strings.ToLower(name) == strings.ToLower(itemRT.Field(i).Name) {
				return itemRV.Field(i), nil
			}
		}

		return NullRv, errors.New("invalid field: [" + name + "]")
	}
}

// Property is 获取属性函数
func Property(name string) func(interface{}) (interface{}, error) {
	fn := PropertyRV(name)
	return func(item interface{}) (interface{}, error) {
		rv, err := fn(item)
		if err != nil {
			return nil, err
		}

		return rv.Interface(), nil
	}
}
