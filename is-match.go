package underscore

// IsMatch is 对象中的属性名与属性值都与map的key和value相同
func IsMatch(item interface{}, properties map[string]interface{}) bool {
	if item == nil || len(properties) == 0 {
		return false
	}

	return All(properties, func(pv interface{}, pn string) bool {
		getValue := Property(pn)
		value, err := getValue(item)
		return err == nil && value == pv
	})
}
