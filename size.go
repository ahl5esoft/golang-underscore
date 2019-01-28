package underscore

// Size is 数组或字典的长度
func Size(source interface{}) int {
	length, _ := parseSource(source)
	return length
}

func (m *query) Size() int {
	return Size(m.Source)
}
