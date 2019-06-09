package underscore

import "reflect"

func (m *query) Uniq(selector interface{}) IQuery {
	if selector == nil {
		selector = func(value, _ interface{}) facade {
			return facade{reflect.ValueOf(value)}
		}
	}

	var mapRV reflect.Value
	var arrRV reflect.Value
	each(m.Source, selector, func(resRV, valueRv, _ reflect.Value) bool {
		if !mapRV.IsValid() {
			mapRT := reflect.MapOf(resRV.Type(), reflect.TypeOf(false))
			mapRV = reflect.MakeMap(mapRT)

			arrRT := reflect.SliceOf(valueRv.Type())
			arrRV = reflect.MakeSlice(arrRT, 0, 0)
		}

		mapValueRV := mapRV.MapIndex(resRV)
		if !mapValueRV.IsValid() {
			mapRV.SetMapIndex(resRV, reflect.ValueOf(true))
			arrRV = reflect.Append(arrRV, valueRv)
		}
		return false
	})

	if mapRV.IsValid() {
		m.Source = arrRV.Interface()
	}
	return m
}

func (m *query) UniqBy(property string) IQuery {
	getPropertyRV := PropertyRV(property)
	return m.Uniq(func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	})
}

func (m enumerable) Uniq(predicate interface{}) IEnumerable {
	return m.Distinct(predicate)
}

func (m enumerable) UniqBy(fieldName string) IEnumerable {
	return m.DistinctBy(fieldName)
}
