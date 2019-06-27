package underscore

// IEnumerable is 迭代器接口
type IEnumerable interface {
	Aggregate(memo interface{}, fn interface{}) IEnumerable
	All(predicate interface{}) bool
	AllBy(dict map[string]interface{}) bool
	Any(predicate interface{}) bool
	AnyBy(dict map[string]interface{}) bool
	Count() int
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
	Group(keySelector interface{}) enumerable
	GroupBy(fieldName string) enumerable
	Index(keySelector interface{}) IEnumerable
	IndexBy(fieldName string) IEnumerable
	Keys() IEnumerable
	Map(selector interface{}) IEnumerable
	MapBy(fieldName string) IEnumerable
	MapMany(selector interface{}) IEnumerable
	MapManyBy(fieldName string) IEnumerable
	Object() IEnumerable
	Reduce(memo interface{}, fn interface{}) IEnumerable
	Select(selector interface{}) IEnumerable
	SelectBy(fieldName string) IEnumerable
	SelectMany(selector interface{}) IEnumerable
	SelectManyBy(fieldName string) IEnumerable
	Size() int
	Skip(count int) IEnumerable
	Take(count int) IEnumerable
	Uniq(selector interface{}) IEnumerable
	UniqBy(fieldName string) IEnumerable
	Value(res interface{})
	Values() IEnumerable
	Where(predicate interface{}) IEnumerable
	WhereBy(dict map[string]interface{}) IEnumerable
}
