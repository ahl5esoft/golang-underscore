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

func (m *query) First() IQuery {
	m.Source = First(m.Source)
	return m
}

func (m enumerable) First() IEnumerable {
	iterator := m.GetEnumerator()
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		return Chain2(
			iterator.GetValue().Interface(),
		)
	}

	return nilEnumerable
}
