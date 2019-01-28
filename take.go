package underscore

// Take is 获取从0开始的n个元素
func Take(source interface{}, count int) interface{} {
	index := 0
	return Where(source, func(_, _ interface{}) bool {
		index = index + 1
		return index <= count
	})
}

func (m *query) Take(count int) IQuery {
	m.Source = Take(m.Source, count)
	return m
}
