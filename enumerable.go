package underscore

type enumerable struct {
	Enumerator func() IEnumerator
}

func (m enumerable) GetEnumerator() IEnumerator {
	return m.Enumerator()
}
