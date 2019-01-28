package underscore

func (m *query) AsParallel() IQuery {
	m.IsParallel = true
	return m
}
