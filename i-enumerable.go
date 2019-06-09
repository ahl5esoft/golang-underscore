package underscore

// IEnumerable is 迭代器接口
type IEnumerable interface {
	All(predicate interface{}) bool
	AllBy(dict map[string]interface{}) bool
	Any(predicate interface{}) bool
	AnyBy(dict map[string]interface{}) bool
	Distinct(selector interface{}) IEnumerable
	DistinctBy(fieldName string) IEnumerable
	Each(action interface{})
	Filter(predicate interface{}) IEnumerable
	FilterBy(dict map[string]interface{}) IEnumerable
	Find(predicate interface{}) IEnumerable
	FindBy(dict map[string]interface{}) IEnumerable
	FindIndex(predicate interface{}) int
	FindIndexBy(dict map[string]interface{}) int
	First() IEnumerable
	GetEnumerator() IEnumerator
	Keys() IEnumerable
	Map(selector interface{}) IEnumerable
	MapBy(fieldName string) IEnumerable
	Object() IEnumerable
	Select(selector interface{}) IEnumerable
	SelectBy(fieldName string) IEnumerable
	Uniq(selector interface{}) IEnumerable
	UniqBy(fieldName string) IEnumerable
	Value(res interface{})
	Values() IEnumerable
	Where(predicate interface{}) IEnumerable
	WhereBy(dict map[string]interface{}) IEnumerable
}
