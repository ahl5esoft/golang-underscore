package underscore

import "sort"

func (m enumerable) Order(selector interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			s := new(sorter)
			s.Sort(
				m.GetEnumerator(),
				selector,
			)
			sort.Sort(s)
			return chainFromValue(s.ValuesValue).GetEnumerator()
		},
	}
}

func (m enumerable) OrderBy(fieldName string) IEnumerable {
	getter := FieldValue(fieldName)
	return m.Order(func(value, _ interface{}) interface{} {
		return getter(value).Interface()
	})
}

func (m enumerable) OrderMany(selectors ...interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			s := new(sorter)
			s.SortMany(
				m.GetEnumerator(),
				selectors...,
			)
			sort.Sort(s)
			return chainFromValue(s.ValuesValue).GetEnumerator()
		},
	}
}

func (m enumerable) OrderManyBy(fieldNames ...string) IEnumerable {
	getters := make([]GetFieldValueFunc, len(fieldNames))
	for i, fieldName := range fieldNames {
		getters[i] = FieldValue(fieldName)
	}

	selectors := make([]interface{}, len(fieldNames))
	for i, getter := range getters {
		selectors[i] = func(value, _ interface{}) interface{} {
			return getter(value).Interface()
		}
	}

	return m.OrderMany(selectors...)
}
