package underscore

// IQuery is interface
type IQuery interface {
	All(interface{}) IQuery
	AllBy(map[string]interface{}) IQuery
	Any(interface{}) IQuery
	AnyBy(map[string]interface{}) IQuery
	AsParallel() IQuery
	Clone() IQuery
	Each(interface{}) IQuery
	Find(interface{}) IQuery
	FindBy(map[string]interface{}) IQuery
	FindIndex(interface{}) IQuery
	FindIndexBy(map[string]interface{}) IQuery
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
	Size() IQuery
	Sort(interface{}) IQuery
	SortBy(string) IQuery
	Take(int) IQuery
	Uniq(interface{}) IQuery
	UniqBy(string) IQuery
	Value() interface{}
	ValueOrDefault(interface{}) interface{}
	Values() IQuery
	Where(interface{}) IQuery
	WhereBy(map[string]interface{}) IQuery
}

// Query is a wrapper
type Query struct {
	isParallel bool
	source     interface{}
}

// Value will return final result
func (q *Query) Value() interface{} {
	return q.source
}

// ValueOrDefault will return final result or default value(if final result is nil)
func (q *Query) ValueOrDefault(defaultValue interface{}) interface{} {
	if q.source == nil {
		return defaultValue
	}

	return q.source
}

// AsParallel will turn on parallel
func (q *Query) AsParallel() IQuery {
	q.isParallel = true
	return q
}

// Chain will cause all future method calls to return wrapped objects
func Chain(source interface{}) IQuery {
	q := new(Query)
	q.source = source
	return q
}
