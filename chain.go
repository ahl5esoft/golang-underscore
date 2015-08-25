package underscore

type Queryer interface {
	Group(func(interface{}, interface{}) (interface{}, error)) Queryer
	GroupBy(string) Queryer
	Index(func(interface{}, interface{}) (interface{}, error)) Queryer
	IndexBy(string) Queryer
	Map(func(interface{}, interface{}) interface{}) Queryer
	Pluck(string) Queryer
	Reduce(func(interface{}, interface{}, interface{}) interface{}, interface{}) Queryer
	Size() Queryer
	Sort(func(interface{},interface{},interface{},interface{}) bool) Queryer
	SortBy(string) Queryer
	Uniq() Queryer
	UniqBy(func(interface{}, int) interface{}) Queryer
	Value() (interface{}, error)
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