package underscore

// Where is 获取所有满足条件
func Where(source, predicate, result interface{}) {
	filter(source, predicate, true, result)
}

// WhereBy is 获取所有满足条件
func WhereBy(source interface{}, properties map[string]interface{}, result interface{}) {
	Where(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	}, result)
}

func (m *query) Where(predicate interface{}) IQuery {
	Where(m.Source, predicate, &m.Source)
	return m
}

func (m *query) WhereBy(properties map[string]interface{}) IQuery {
	WhereBy(m.Source, properties, &m.Source)
	return m
}
