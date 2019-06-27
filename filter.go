package underscore

func (m enumerable) Filter(predicate interface{}) IEnumerable {
	return m.Where(predicate)
}

func (m enumerable) FilterBy(dict map[string]interface{}) IEnumerable {
	return m.WhereBy(dict)
}
