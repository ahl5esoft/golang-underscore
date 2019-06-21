package underscore

func (m enumerable) Count() int {
	iterator := m.GetEnumerator()
	count := 0
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		count++
	}

	return count
}
