package underscore

type Queryer interface {
	All(interface{}) Queryer
	AllBy(map[string]interface{}) Queryer
	Any(interface{}) Queryer
	AnyBy(map[string]interface{}) Queryer
	Clone() Queryer
	Each(interface{}) Queryer
	Find(interface{}) Queryer
	FindBy(map[string]interface{}) Queryer
	First() Queryer
	Group(interface{}) Queryer
	GroupBy(string) Queryer
	Index(interface{}) Queryer
	IndexBy(string) Queryer
	Keys() Queryer
	Map(interface{}) Queryer
	Pluck(string) Queryer
	Reduce(interface{}, interface{}) Queryer
	Select(interface{}) Queryer
	SelectBy(map[string]interface{}) Queryer
	Size() Queryer
	Sort(interface{}) Queryer
	SortBy(string) Queryer
	Take(int) Queryer
	Uniq(interface{}) Queryer
	UniqBy(string) Queryer
	Value() (interface{}, error)
	Values() Queryer
}

type Query struct {
	err error
	source interface{}
}

func (this *Query) Value() (interface{}, error) {
	return this.source, this.err
}

func Chain(source interface{}) Queryer {
	return &Query{ nil, source };
}