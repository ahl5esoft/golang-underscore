package underscore

type Queryer interface {
	Group(func(item interface{}) interface{}) Queryer
	GroupBy(string) Queryer
	Index(func(item interface{}) interface{}) Queryer
	IndexBy(string) Queryer
	Count() Queryer
	Value() interface{}
}

type Query struct {
	source interface{}
}

func (this *Query) Value() interface{} {
	return this.source
}

func Chain(source interface{}) Queryer {
	return &Query{ source };
}