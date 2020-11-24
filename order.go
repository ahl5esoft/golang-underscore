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
	return m.Order(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
