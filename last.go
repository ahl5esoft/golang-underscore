package underscore

import "reflect"

func (m enumerable) Last() IEnumerable {
	iterator := m.GetEnumerator()
	var lastValue reflect.Value
	for ok := iterator.MoveNext(); ok; ok = iterator.MoveNext() {
		lastValue = iterator.GetValue()
	}

	if lastValue.IsValid() {
		return chainFromValue(lastValue)
	}

	return nilEnumerable
}
