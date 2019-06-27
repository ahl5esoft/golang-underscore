package underscore

// IsMatch is 对象中的属性名与属性值都与map的key和value相同
func IsMatch(item interface{}, properties map[string]interface{}) bool {
	if item == nil || len(properties) == 0 {
		return false
	}

	return Chain(properties).All(func(v interface{}, k string) bool {
		return Property(k)(item) == v
	})
}
