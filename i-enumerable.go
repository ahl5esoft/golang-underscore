package underscore

// IEnumerable is 迭代器接口
type IEnumerable interface {
	Aggregate(fn interface{}, memo interface{}) IEnumerable
	All(predicate interface{}) bool
	AllBy(fields map[string]interface{}) bool
	Any(predicate interface{}) bool
	AnyBy(fields map[string]interface{}) bool
	Count() int
	Distinct(selector interface{}) IEnumerable
	DistinctBy(fieldName string) IEnumerable
	Each(action interface{})
	Except(predicate interface{}) IEnumerable
	ExceptBy(fields map[string]interface{}) IEnumerable
	Filter(predicate interface{}) IEnumerable
	FilterBy(fields map[string]interface{}) IEnumerable
	Find(predicate interface{}) IEnumerable
	FindBy(fields map[string]interface{}) IEnumerable
	FindIndex(predicate interface{}) int
	FindIndexBy(fields map[string]interface{}) int
	First() IEnumerable
	GetEnumerator() IEnumerator
	Group(keySelector interface{}) enumerable
	GroupBy(fieldName string) enumerable
	Index(keySelector interface{}) IEnumerable
	IndexBy(fieldName string) IEnumerable
	Keys() IEnumerable
	Last() IEnumerable
	Map(selector interface{}) IEnumerable
	MapBy(fieldName string) IEnumerable
	MapMany(selector interface{}) IEnumerable
	MapManyBy(fieldName string) IEnumerable
	Object() IEnumerable
	Order(selector interface{}) IEnumerable
	OrderBy(fieldName string) IEnumerable
	Reduce(fn interface{}, memo interface{}) IEnumerable
	Reject(predicate interface{}) IEnumerable
	RejectBy(fields map[string]interface{}) IEnumerable
	Reverse(selector interface{}) IEnumerable
	ReverseBy(fieldName string) IEnumerable
	Select(selector interface{}) IEnumerable
	SelectBy(fieldName string) IEnumerable
	SelectMany(selector interface{}) IEnumerable
	SelectManyBy(fieldName string) IEnumerable
	Sort(selector interface{}) IEnumerable
	SortBy(fieldName string) IEnumerable
	Size() int
	Skip(count int) IEnumerable
	Take(count int) IEnumerable
	Uniq(selector interface{}) IEnumerable
	UniqBy(fieldName string) IEnumerable
	Value(res interface{})
	Values() IEnumerable
	Where(predicate interface{}) IEnumerable
	WhereBy(fields map[string]interface{}) IEnumerable
}
