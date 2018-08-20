package underscore

// Queryer is interface
type Queryer interface {
	All(interface{}) Queryer
	AllBy(map[string]interface{}) Queryer
	Any(interface{}) Queryer
	AnyBy(map[string]interface{}) Queryer
	AsParallel() Queryer
	Clone() Queryer
	Each(interface{}) Queryer
	Find(interface{}) Queryer
	FindBy(map[string]interface{}) Queryer
	FindIndex(interface{}) Queryer
	FindIndexBy(map[string]interface{}) Queryer
	First() Queryer
	Group(interface{}) Queryer
	GroupBy(string) Queryer
	Index(interface{}) Queryer
	IndexBy(string) Queryer
	Keys() Queryer
	Last() Queryer
	Map(interface{}) Queryer
	MapBy(string) Queryer
	Object() Queryer
	Pluck(string) Queryer
	Range(int, int, int) Queryer
	Reduce(interface{}, interface{}) Queryer
	Reject(interface{}) Queryer
	RejectBy(map[string]interface{}) Queryer
	Size() Queryer
	Sort(interface{}) Queryer
	SortBy(string) Queryer
	Take(int) Queryer
	Uniq(interface{}) Queryer
	UniqBy(string) Queryer
	Value() interface{}
	Values() Queryer
	Where(interface{}) Queryer
	WhereBy(map[string]interface{}) Queryer
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

// AsParallel will turn on parallel
func (q *Query) AsParallel() Queryer {
	q.isParallel = true
	return q
}

// Chain will cause all future method calls to return wrapped objects
func Chain(source interface{}) Queryer {
	q := new(Query)
	q.source = source
	return q
}
