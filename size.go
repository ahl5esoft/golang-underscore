package underscore

// Size is 数组或字典的长度
func Size(source interface{}) int {
	length, _ := parseSource(source)
	return length
}

// Size is IQuery's Method
func (q *Query) Size() IQuery {
	q.source = Size(q.source)
	return q
}
