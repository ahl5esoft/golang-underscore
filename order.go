package underscore

import "sort"

func (m enumerable) Order(selectors ...interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			s := new(sorter)
			s.SortMultiple(
				m.GetEnumerator(),
				selectors,
			)
			sort.Sort(s)
			return chainFromValue(s.ValuesValue).GetEnumerator()
		},
	}
}

func (m enumerable) OrderBy(fieldName string) IEnumerable {
	getter := FieldValue(fieldName)
	return m.Order(func(value, _ interface{}) interface{} {
		return getter(value)
	})
}

func (m enumerable) OrderByFields(fieldNames ...string) IEnumerable {
	var selectors []interface{}
	for _, fieldName := range fieldNames {
		getter := FieldValue(fieldName)
		selectors = append(selectors, func(value, _ interface{}) interface{} {
			return getter(value)
		})
	}
	return m.Order(selectors...)
}
