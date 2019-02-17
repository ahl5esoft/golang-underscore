package underscore

// Reject is 排除
func Reject(source, predicate, result interface{}) {
	filter(source, predicate, false, result)
}

// RejectBy is 根据属性排除
func RejectBy(source interface{}, properties map[string]interface{}, result interface{}) {
	Reject(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	}, result)
}

func (m *query) Reject(predicate interface{}) IQuery {
	Reject(m.Source, predicate, &m.Source)
	return m
}

func (m *query) RejectBy(properties map[string]interface{}) IQuery {
	RejectBy(m.Source, properties, &m.Source)
	return m
}
