package underscore

// IEnumerable is 迭代器接口
type IEnumerable interface {
	All(predicate interface{}) bool
	AllBy(dict map[string]interface{}) bool
	Any(predicate interface{}) bool
	AnyBy(dict map[string]interface{}) bool
	Each(action interface{})
	Filter(predicate interface{}) IEnumerable
	FilterBy(dict map[string]interface{}) IEnumerable
	Find(predicate interface{}) IEnumerable
	FindBy(dict map[string]interface{}) IEnumerable
	First() IEnumerable
	GetEnumerator() IEnumerator
	Value(res interface{})
	Where(predicate interface{}) IEnumerable
	WhereBy(dict map[string]interface{}) IEnumerable
}
