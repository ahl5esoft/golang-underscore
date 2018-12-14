package underscore

// Pluck is 从source中取出所有property
func Pluck(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Map(source, func(value, _ interface{}) Facade {
		return Facade{
			getPropertyRV(value),
		}
	})
}

// Pluck is IQuery's method
func (q *Query) Pluck(property string) IQuery {
	q.source = Pluck(q.source, property)
	return q
}
