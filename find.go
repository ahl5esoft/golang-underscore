package underscore

import (
	"reflect"
)

// Find is 根据断言获取元素
func Find(source, predicate interface{}) interface{} {
	var ok bool
	var matcher interface{}
	each(source, predicate, func(resRV, valueRV, _ reflect.Value) bool {
		ok = resRV.Bool()
		if ok {
			matcher = valueRV.Interface()
		}
		return ok
	})
	return matcher
}

// FindBy is 根据属性值获取元素
func FindBy(source interface{}, properties map[string]interface{}) interface{} {
	return Find(source, func(value, _ interface{}) bool {
		return IsMatch(value, properties)
	})
}

// Find is IQuery's method
func (q *Query) Find(predicate interface{}) IQuery {
	q.source = Find(q.source, predicate)
	return q
}

// FindBy is IQuery's method
func (q *Query) FindBy(properties map[string]interface{}) IQuery {
	q.source = FindBy(q.source, properties)
	return q
}
