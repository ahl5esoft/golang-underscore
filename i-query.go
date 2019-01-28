package underscore

// IQuery is interface
type IQuery interface {
	All(interface{}) bool
	AllBy(map[string]interface{}) bool
	Any(interface{}) bool
	AnyBy(map[string]interface{}) bool
	AsParallel() IQuery
	Clone() IQuery
	Each(interface{}) IQuery
	Find(interface{}) IQuery
	FindBy(map[string]interface{}) IQuery
	FindIndex(interface{}) int
	FindIndexBy(map[string]interface{}) int
	FindLastIndex() IQuery
	First() IQuery
	Group(interface{}) IQuery
	GroupBy(string) IQuery
	Index(interface{}) IQuery
	IndexBy(string) IQuery
	Keys() IQuery
	Last() IQuery
	Map(interface{}) IQuery
	Object() IQuery
	Pluck(string) IQuery
	Range(int, int, int) IQuery
	Reduce(interface{}, interface{}) IQuery
	Reject(interface{}) IQuery
	RejectBy(map[string]interface{}) IQuery
	Size() int
	Sort(interface{}) IQuery
	SortBy(string) IQuery
	Take(int) IQuery
	Uniq(interface{}) IQuery
	UniqBy(string) IQuery
	Value(v interface{})
	ValueOrDefault(interface{}) interface{}
	Values() IQuery
	Where(interface{}) IQuery
	WhereBy(map[string]interface{}) IQuery
}
