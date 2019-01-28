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

func (m *query) Where(predicate interface{}) IQuery {
	m.Source = Where(m.Source, predicate)
	return m
}

func (m *query) WhereBy(properties map[string]interface{}) IQuery {
	m.Source = WhereBy(m.Source, properties)
	return m
}
