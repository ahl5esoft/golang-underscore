package underscore

func (m enumerable) Except(predicate interface{}) IEnumerable {
	return m.Reject(predicate)
}

func (m enumerable) ExceptBy(dict map[string]interface{}) IEnumerable {
	return m.RejectBy(dict)
}
