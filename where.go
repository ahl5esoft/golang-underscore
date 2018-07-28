package underscore

// Where is 获取所有满足条件
func Where(source, predicate interface{}) interface{} {
	return filter(source, predicate, true)
}

// WhereBy is 获取所有满足条件
func WhereBy(source interface{}, properties map[string]interface{}) interface{} {
	return Where(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

// Where is Queryer's method
func (q *Query) Where(predicate interface{}) Queryer {
	q.source = Where(q.source, predicate)
	return q
}

// WhereBy is Queryer's method
func (q *Query) WhereBy(properties map[string]interface{}) Queryer {
	q.source = WhereBy(q.source, properties)
	return q
}
