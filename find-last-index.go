package underscore

import "reflect"

// FindLastIndex gets the last index of the argument
func FindLastIndex(source interface{}) int {
	lastIndex := -1

	if !IsArray(source) {
		return lastIndex
	}

	sourceLength := reflect.ValueOf(source).Len()
	lastIndex = sourceLength - 1

	return lastIndex
}

// FindLastIndex is Queryer's method
func (q *Query) FindLastIndex() Queryer {
	q.source = FindLastIndex(q.source)
	return q
}
