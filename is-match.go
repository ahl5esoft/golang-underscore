package underscore

// IsMatch is 对象中的属性名与属性值都与map的key和value相同
func IsMatch(item interface{}, fields map[string]interface{}) bool {
	if item == nil || len(fields) == 0 {
		return false
	}

	return Chain(fields).All(func(v interface{}, k string) bool {
		return Field(k)(item) == v
	})
}
