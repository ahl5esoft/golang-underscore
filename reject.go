package underscore

// Reject is 排除
func Reject(source, predicate interface{}) interface{} {
	return filter(source, predicate, false)
}

// RejectBy is 根据属性排除
func RejectBy(source interface{}, properties map[string]interface{}) interface{} {
	return Reject(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

// Reject is IQuery's Method
func (q *Query) Reject(predicate interface{}) IQuery {
	q.source = Reject(q.source, predicate)
	return q
}

// RejectBy is IQuery's Method
func (q *Query) RejectBy(properties map[string]interface{}) IQuery {
	q.source = RejectBy(q.source, properties)
	return q
}
