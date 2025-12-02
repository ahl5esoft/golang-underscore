package underscore

import "sort"

func (m enumerable) Order(selectors ...interface{}) IEnumerable {
	return enumerable{
		Enumerator: func() IEnumerator {
			s := new(sorter)
			s.Sort(
				m.GetEnumerator(),
				selectors...,
			)
			sort.Sort(s)
			return chainFromValue(s.ValuesValue).GetEnumerator()
		},
	}
}

func (m enumerable) OrderBy(fieldNames ...string) IEnumerable {
	selectors := make([]interface{}, len(fieldNames))
	for i, fieldName := range fieldNames {
		getter := FieldValue(fieldName)
		selectors[i] = func(value, _ interface{}) facade {
			return facade{
				getter(value),
			}
		}
	}
	return m.Order(selectors...)
}
