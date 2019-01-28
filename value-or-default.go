package underscore

func (m *query) ValueOrDefault(defaultValue interface{}) interface{} {
	if m.Source == nil {
		return defaultValue
	}

	return m.Source
}
