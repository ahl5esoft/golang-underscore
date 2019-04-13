package underscore

import (
	"reflect"
)

// MapMany is 将序列的每个元素投影到一个序列，并将结果序列合并为一个序列
func MapMany(source, selector interface{}) interface{} {
	var sliceRV reflect.Value
	each(source, selector, func(resRV, valueRV, _ reflect.Value) bool {
		if !sliceRV.IsValid() {
			if !(resRV.Kind() == reflect.Slice || resRV.Kind() == reflect.Array) {
				panic("selector的返回值必须是Array或Slice")
			}

			sliceRT := reflect.SliceOf(resRV.Type().Elem())
			sliceRV = reflect.MakeSlice(sliceRT, 0, 0)
		}

		if resRV.Kind() == reflect.Slice {
			sliceRV = reflect.AppendSlice(sliceRV, resRV)
		} else {
			for i := 0; i < resRV.Len(); i++ {
				sliceRV = reflect.Append(sliceRV, resRV.Index(i))
			}
		}
		return false
	})
	if sliceRV.IsValid() {
		return sliceRV.Interface()
	}

	return nil
}

// MapManyBy is 将序列的每个元素的property值(值必须是一个序列)合并为一个序列
func MapManyBy(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return MapMany(source, func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	})
}

func (m *query) MapMany(selector interface{}) IQuery {
	m.Source = MapMany(m.Source, selector)
	return m
}

func (m *query) MapManyBy(property string) IQuery {
	m.Source = MapManyBy(m.Source, property)
	return m
}
