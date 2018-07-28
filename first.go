package underscore

// First is 获取第一个元素
func First(source interface{}) interface{} {
	length, getKeyValue := parseSource(source)
	if length == 0 {
		return nil
	}

	valueRV, _ := getKeyValue(0)
	return valueRV.Interface()
}

// First is Queryer's method
func (q *Query) First() Queryer {
	q.source = First(q.source)
	return q
}
