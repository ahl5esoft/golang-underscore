package underscore

type Queryer interface {
	Group(func(interface{}) (interface{}, error)) Queryer
	GroupBy(string) Queryer
	Index(func(interface{}) (interface{}, error)) Queryer
	IndexBy(string) Queryer
	Map(func(interface{}) interface{}) Queryer
	Uniq() Queryer
	UniqBy(func(interface{}) interface{}) Queryer
	Count() Queryer
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