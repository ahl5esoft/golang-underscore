package underscore

func Values(source interface{}) (interface{}, error) {
	return mapFromEach(source, 0)
}

//Chain
func (this *Query) Values() Queryer {
	if this.err == nil {
		this.source, this.err = Values(this.source)
	}
	return this
}