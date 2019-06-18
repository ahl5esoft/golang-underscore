package underscore

import "reflect"

func (m *query) Group(keySelector interface{}) IQuery {
	var groupRV reflect.Value
	each(m.Source, keySelector, func(groupKeyRV, valueRV, _ reflect.Value) bool {
		groupValueRT := reflect.SliceOf(valueRV.Type())
		if !groupRV.IsValid() {
			groupRT := reflect.MapOf(groupKeyRV.Type(), groupValueRT)
			groupRV = reflect.MakeMap(groupRT)
		}

		valuesRV := groupRV.MapIndex(groupKeyRV)
		if !valuesRV.IsValid() {
			valuesRV = reflect.MakeSlice(groupValueRT, 0, 0)
		}
		valuesRV = reflect.Append(valuesRV, valueRV)

		groupRV.SetMapIndex(groupKeyRV, valuesRV)
		return false
	})
	if groupRV.IsValid() {
		m.Source = groupRV.Interface()
	}

	return m
}

func (m *query) GroupBy(property string) IQuery {
	getPropertyRV := PropertyRV(property)
	return m.Group(func(value, _ interface{}) facade {
		return facade{
			getPropertyRV(value),
		}
	})
}

func (m enumerable) Group(keySelector interface{}) enumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			groupRVs := make(map[interface{}]reflect.Value)
			iterator := m.GetEnumerator()
			keySelectorRV := reflect.ValueOf(keySelector)
			keyRVs := make([]reflect.Value, 0)
			for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
				keyRV := getFuncReturnRV(keySelectorRV, iterator)
				key := keyRV.Interface()
				groupRV, ok := groupRVs[key]
				if !ok {
					groupRV = reflect.MakeSlice(
						reflect.SliceOf(
							iterator.GetValue().Type(),
						),
						0,
						0,
					)
					keyRVs = append(keyRVs, keyRV)
				}
				groupRVs[key] = reflect.Append(
					groupRV,
					iterator.GetValue(),
				)
			}
			index := 0
			return &enumerator{
				MoveNextFunc: func() (valueRV reflect.Value, keyRV reflect.Value, ok bool) {
					if ok = index < len(keyRVs); ok {
						keyRV = keyRVs[index]
						valueRV = groupRVs[keyRV.Interface()]
						index++
					}
					return
				},
			}
		},
	}
}

func (m enumerable) GroupBy(fieldName string) enumerable {
	getter := PropertyRV(fieldName)
	return m.Group(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
