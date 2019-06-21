package underscore

// IQuery is interface
type IQuery interface {
	All(interface{}) bool
	AllBy(map[string]interface{}) bool
	Any(interface{}) bool
	AnyBy(map[string]interface{}) bool
	AsParallel() IQuery
	Each(interface{})
	Find(interface{}) IQuery
	FindBy(map[string]interface{}) IQuery
	FindIndex(interface{}) int
	FindIndexBy(map[string]interface{}) int
	First() IQuery
	Group(interface{}) IQuery
	GroupBy(string) IQuery
	Index(interface{}) IQuery
	IndexBy(string) IQuery
	Keys() IQuery
	Last() IQuery
	Map(interface{}) IQuery
	MapBy(string) IQuery
	MapMany(interface{}) IQuery
	MapManyBy(string) IQuery
	Object() IQuery
	Reduce(interface{}, interface{}) IQuery
	Reject(interface{}) IQuery
	RejectBy(map[string]interface{}) IQuery
	Reverse(interface{}) IQuery
	ReverseBy(string) IQuery
	Size() int
	Sort(interface{}) IQuery
	SortBy(string) IQuery
	Take(int) IQuery
	Uniq(interface{}) IQuery
	UniqBy(string) IQuery
	Value(v interface{})
	Values() IQuery
	Where(interface{}) IQuery
	WhereBy(map[string]interface{}) IQuery
}
