package underscore

func (m enumerable) Map(selector interface{}) IEnumerable {
	return m.Select(selector)
}

func (m enumerable) MapBy(fieldName string) IEnumerable {
	return m.SelectBy(fieldName)
}
