package underscore

// Take is 获取从0开始的n个元素
func Take(source interface{}, count int) interface{} {
	index := 0
	return Where(source, func(_, _ interface{}) bool {
		index = index + 1
		return index <= count
	})
}

// Take is Query's method
func (q *Query) Take(count int) IQuery {
	q.source = Take(q.source, count)
	return q
}
