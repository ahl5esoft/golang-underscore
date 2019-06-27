package underscore

func (m enumerable) Reduce(memo interface{}, fn interface{}) IEnumerable {
	return m.Aggregate(memo, fn)
}
