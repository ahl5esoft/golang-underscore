package underscore

func (m enumerable) First() IEnumerable {
	iterator := m.GetEnumerator()
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		return chainFromRV(
			iterator.GetValue(),
		)
	}

	return nilEnumerable
}
