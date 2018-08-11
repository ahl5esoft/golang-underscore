package underscore

// Pluck is 从source中取出所有property
func Pluck(source interface{}, property string) interface{} {
	getPropertyRV := PropertyRV(property)
	return Map(source, func(value, _ interface{}) Facade {
		rv, _ := getPropertyRV(value)
		return Facade{rv}
	})
}

// Pluck is Queryer's method
func (q *Query) Pluck(property string) Queryer {
	q.source = Pluck(q.source, property)
	return q
}
