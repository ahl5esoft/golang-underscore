package underscore

// FindLastIndex gets the last index of the argument
func FindLastIndex(source interface{}) int {
	if !IsArray(source) {
		return -1
	}

	return Size(source) - 1
}

func (m *query) FindLastIndex() IQuery {
	m.Source = FindLastIndex(m.Source)
	return m
}
