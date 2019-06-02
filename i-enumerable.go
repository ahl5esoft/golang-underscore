package underscore

// IEnumerable is 迭代器接口
type IEnumerable interface {
	All(predicate interface{}) bool
	AllBy(dict map[string]interface{}) bool
	Any(predicate interface{}) bool
	AnyBy(dict map[string]interface{}) bool
	Find(predicate interface{}) IEnumerable
	FindBy(dict map[string]interface{}) IEnumerable
	First() IEnumerable
	GetEnumerator() IEnumerator
	Value(res interface{})
}
