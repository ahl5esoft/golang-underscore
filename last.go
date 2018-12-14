package underscore

// Last is 最后元素
func Last(source interface{}) interface{} {
	length, getKeyValue := parseSource(source)
	if length == 0 {
		return nil
	}

	valueRV, _ := getKeyValue(length - 1)
	return valueRV.Interface()
}

// Last is IQuery's method
func (q *Query) Last() IQuery {
	q.source = Last(q.source)
	return q
}
