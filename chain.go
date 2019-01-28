package underscore

// Chain will cause all future method calls to return wrapped objects
func Chain(source interface{}) IQuery {
	return &query{
		Source: source,
	}
}
