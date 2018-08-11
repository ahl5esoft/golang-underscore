package underscore

// Size is 数组或字典的长度
func Size(source interface{}) int {
	length, _ := parseSource(source)
	return length
}

// Size is Queryer's Method
func (q *Query) Size() Queryer {
	q.source = Size(q.source)
	return q
}
