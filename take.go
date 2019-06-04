package underscore

func (m *query) Take(count int) IQuery {
	index := 0
	return m.Where(func(_, _ interface{}) bool {
		index = index + 1
		return index <= count
	})
}
