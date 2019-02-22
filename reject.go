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

func (m *query) Reject(predicate interface{}) IQuery {
	m.Source = Reject(m.Source, predicate)
	return m
}

func (m *query) RejectBy(properties map[string]interface{}) IQuery {
	m.Source = RejectBy(m.Source, properties)
	return m
}
