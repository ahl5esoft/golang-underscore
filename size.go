package underscore

func (m *query) Size() int {
	length, _ := parseSource(m.Source)
	return length
}

func (m enumerable) Size() int {
	return m.Count()
}
