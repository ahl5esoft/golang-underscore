package underscore

// FindLastIndex gets the last index of the argument
func FindLastIndex(source interface{}) int {
	if !IsArray(source) {
		return -1
	}

	return Size(source) - 1
}

// FindLastIndex is Queryer's method
func (q *Query) FindLastIndex() IQuery {
	q.source = FindLastIndex(q.source)
	return q
}
