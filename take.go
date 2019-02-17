package underscore

// Take is 获取从0开始的n个元素
func Take(source interface{}, count int, result interface{}) {
	index := 0
	Where(source, func(_, _ interface{}) bool {
		index = index + 1
		return index <= count
	}, result)
}

func (m *query) Take(count int) IQuery {
	Take(m.Source, count, &m.Source)
	return m
}
