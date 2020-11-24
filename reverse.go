package underscore

func (m enumerable) Reverse(selector interface{}) IEnumerable {
	return m.Sort(selector).Sort(func(_ interface{}, i int) int {
		return -i
	})
}

func (m enumerable) ReverseBy(fieldName string) IEnumerable {
	getter := FieldValue(fieldName)
	return m.Reverse(func(value, _ interface{}) facade {
		return facade{
			getter(value),
		}
	})
}
