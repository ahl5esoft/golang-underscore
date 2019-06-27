package underscore

func (m enumerable) Uniq(predicate interface{}) IEnumerable {
	return m.Distinct(predicate)
}

func (m enumerable) UniqBy(fieldName string) IEnumerable {
	return m.DistinctBy(fieldName)
}
