package underscore

func (m enumerable) First() IEnumerable {
	iterator := m.GetEnumerator()
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		return chainFromValue(
			iterator.GetValue(),
		)
	}

	return nilEnumerable
}
