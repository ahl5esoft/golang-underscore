package underscore

func (m enumerable) Sort(selector interface{}) IEnumerable {
	return m.Order(selector)
}

func (m enumerable) SortBy(fieldName string) IEnumerable {
	return m.OrderBy(fieldName)
}
