package underscore

func (m enumerable) MapMany(selector interface{}) IEnumerable {
	return m.SelectMany(selector)
}

func (m enumerable) MapManyBy(fieldName string) IEnumerable {
	return m.SelectManyBy(fieldName)
}
