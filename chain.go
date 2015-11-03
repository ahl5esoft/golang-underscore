package underscore

type Queryer interface {
	All(func(interface{}, interface{}) (bool, error)) Queryer
	AllBy(map[string]interface{}) Queryer
	Any(func(interface{}, interface{}) (bool, error)) Queryer
	AnyBy(map[string]interface{}) Queryer
	Clone() Queryer
	Find(func(interface{}, interface{}) (bool, error)) Queryer
	FindBy(map[string]interface{}) Queryer
	Group(func(interface{}, interface{}) (interface{}, error)) Queryer
	GroupBy(string) Queryer
	Index(func(interface{}, interface{}) (interface{}, error)) Queryer
	IndexBy(string) Queryer
	Map(func(interface{}, interface{}) (interface{}, error)) Queryer
	Pluck(string) Queryer
	Reduce(func(interface{}, interface{}, interface{}) (interface{}, error), interface{}) Queryer
	Select(func(interface{}, interface{}) (bool, error)) Queryer
	SelectBy(map[string]interface{}) Queryer
	Size() Queryer
	Sort(func(interface{}, interface{}, interface{}, interface{}) bool) Queryer
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